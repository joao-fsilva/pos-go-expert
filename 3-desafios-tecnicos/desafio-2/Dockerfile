FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o stress-test main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/stress-test .

ENTRYPOINT ["/app/stress-test"]