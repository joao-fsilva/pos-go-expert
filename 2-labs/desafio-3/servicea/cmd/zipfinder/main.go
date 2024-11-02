package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"time"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func validateZipCode(cep string) bool {
	var cepRegex = regexp.MustCompile(`^\d{8}$`)
	return cepRegex.MatchString(cep)
}

type RequestParams struct {
	CEP string `json:"cep"`
}

func initProvider() (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("servicea"),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "otel-collector:4317",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		var reqParams RequestParams
		if err := json.NewDecoder(r.Body).Decode(&reqParams); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if !validateZipCode(reqParams.CEP) {
			http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		shutdown, err := initProvider()
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if err := shutdown(ctx); err != nil {
				log.Fatal("failed to shutdown TracerProvider: %w", err)
			}
		}()

		tracer := otel.Tracer("microservice-tracer")

		carrier := propagation.HeaderCarrier(r.Header)
		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

		ctx, span := tracer.Start(ctx, "chamada serviceb")
		defer span.End()

		time.Sleep(time.Millisecond * 3)

		url := os.Getenv("SERVICEB_URL") + "weather?zipcode=" + reqParams.CEP

		var req *http.Request
		req, err = http.NewRequestWithContext(ctx, "GET", url, nil)

		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		res, err := http.DefaultClient.Do(req)
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Printf("Error: %v", err)
			writeErrorResponse(w, http.StatusInternalServerError, "Error processing the request")
			return
		}

		if err != nil {
			log.Printf("Error: %v", err)
			writeErrorResponse(w, http.StatusInternalServerError, "Error processing the request")
			return
		}

		body, err := io.ReadAll(res.Body)

		var data map[string]interface{}
		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Printf("Error: %v", err)
			writeErrorResponse(w, http.StatusInternalServerError, "Error processing the request")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8081", nil)
}
