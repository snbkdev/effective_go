package main

import (
	"context"
	"time"
)

func AsyncResult(ctx context.Context) {
	requestsCh := make(chan request)
	updateTicker := time.NewTicker(30 * time.Second)

	var updateResultCh chan data

	for {
		select {
		case <- updateTicker.C:
			updateResultCh = make(chan data, 1)
			go loadUpdate(updateResultCh)

		case updateResult := <- updateResultCh:
			processUpdate(updateResult)

			updateResultCh = nil

		case request := <- requestsCh:
			processRequest(request)

		case <- ctx.Done():
			return
		}
	}
}

func loadUpdate(responseCh chan data) {}