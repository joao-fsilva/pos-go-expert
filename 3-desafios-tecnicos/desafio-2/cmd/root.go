package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	url         string
	requests    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "A CLI to make multiple HTTP requests",
	Run: func(cmd *cobra.Command, args []string) {
		runRequests(url, requests, concurrency)
	},
}

func init() {
	rootCmd.Flags().StringVar(&url, "url", "http://example.com", "URL to send requests to")
	rootCmd.Flags().IntVar(&requests, "requests", 100, "Number of requests to send")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Number of concurrent requests")
}

func runRequests(url string, requests int, concurrency int) {
	startTime := time.Now()

	ch := make(chan bool, concurrency)

	var totalRequests, total200, totalOther int
	statusCodes := make(map[int]int)

	var mu sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < requests; i++ {
		wg.Add(1)

		ch <- true

		go func(i int) {
			defer wg.Done()

			log.Printf("Realizando requisição %d para a URL %s\n", i+1, url)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Erro na requisição %d: %s\n", i+1, err)
				<-ch
				return
			}
			defer resp.Body.Close()

			mu.Lock()

			totalRequests++

			if resp.StatusCode == 200 {
				total200++
			} else {
				totalOther++
				statusCodes[resp.StatusCode]++
			}
			mu.Unlock()

			log.Printf("Requisição %d concluída com status: %d\n", i+1, resp.StatusCode)
			<-ch
		}(i)
	}

	wg.Wait()

	elapsedTime := time.Since(startTime)

	totalSeconds := elapsedTime.Seconds()

	fmt.Printf("\nRelatório de Execução:\n")
	fmt.Printf("Tempo total gasto na execução: %.1f segundos\n", totalSeconds)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", total200)

	if totalOther > 0 {
		for code, count := range statusCodes {
			fmt.Printf("  Status %d: %d requisições\n", code, count)
		}
	} else {
		fmt.Println("Nenhuma requisição com status diferente de 200.")
	}
}

func Execute() error {
	return rootCmd.Execute()
}
