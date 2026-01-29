package main

import (
	"context"
	"time"
)

func BlockingCaseExample(ctx context.Context) {
	requestsCh := make(chan request)
	updatesCh := make(chan data)
	updateTicker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <- updateTicker.C:
			newData := <- updatesCh
			processUpdate(newData)

		case request := <- requestsCh:
			processRequest(request)
		
		case <- ctx.Done():
			return
		}
	}
}

type request struct {
	//
}

type data struct {
	//
}

func processUpdate(_ data) {
	//
}

func processRequest(_ request) {
	//
}