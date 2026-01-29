package main

import (
	"context"
	"sync"
	"time"
)

var (
	cache = []string{}

	cacheMutex = &sync.Mutex{}
)

func useTheData(cacheCopy []string) {
	//
}

func loadUpdates() []string {
	return []string{}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
	defer cancel()

	go updater(ctx)
	go reader(ctx)
}

func reader(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			return
		default:
			//
		}

		cacheMutex.Lock()
		cacheCopy := make([]string, len(cache))
		copy(cacheCopy, cache)
		cacheMutex.Unlock()

		useTheData(cacheCopy)
	}
}

func updater(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			updates := loadUpdates()

			cacheMutex.Lock()
			cache = updates
			cacheMutex.Unlock()

		case <-ctx.Done():
			return
		}
	}
}