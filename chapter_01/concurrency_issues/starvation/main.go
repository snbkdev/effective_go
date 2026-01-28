package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	sites = map[string]int{
		"https://www.coreyscott.dev": http.StatusOK,
		"https://golang.org/":        http.StatusOK,
	}
	sitesMutex = &sync.Mutex{}
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go selfishChecker(ctx, wg, "https://www.coreyscott.dev")

	wg.Add(1)
	go politeChecker(ctx, wg, "https://golang.org")

	wg.Wait()
}

func selfishChecker(ctx context.Context, wg *sync.WaitGroup, url string) {
	defer wg.Done()

	totalAttempts := 0

	for{
		totalAttempts++

		select{
		case <- ctx.Done():
			fmt.Printf("selfish: total updates %d\n", totalAttempts)
			return
		default:
			// continue
		}

		sitesMutex.Lock()

		resp, err := http.Head(url)
		if err != nil {
			sitesMutex.Unlock()
			continue
		}

		sites[url] = resp.StatusCode

		sitesMutex.Unlock()
	}
}

func politeChecker(ctx context.Context, wg *sync.WaitGroup, url string) {
	defer wg.Done()

	totalAttempts := 0

	for{
		totalAttempts++

		select{
		case <- ctx.Done():
			fmt.Printf("polite: total updates %d\n", totalAttempts)
			return
		default:
			// continue
		}

		resp, err := http.Head(url)
		if err != nil {
			continue
		}

		sitesMutex.Lock()

		sites[url] = resp.StatusCode

		sitesMutex.Unlock()
	}
}