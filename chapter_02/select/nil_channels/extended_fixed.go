package main

import (
	"context"
	"time"
)

func extendedExampleFixed(ctx context.Context) {
	var updateCh chan string

	updateTicker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <- ctx.Done():
			//
			return

		case <- updateTicker.C:
			updateCh = performUpdateAsync()

		case result := <- updateCh:
			useResult(result)
			updateCh = nil
		}
	}
}

func performUpdateAsync() chan string {
	result := make(chan string, 1)

	go func() {
		defer close(result)
	}()

	return result
}