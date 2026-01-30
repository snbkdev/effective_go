package main

import (
	"context"
	"time"
)

func extendedExample(ctx context.Context) {
	updateTicker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <- ctx.Done():
			//
			return

		case <- updateTicker.C:
			result := performUpdate()
			useResult(result)
		}
	}
}

func performUpdate() string {
	//
	return ""
}

func useResult(in string) {
	//
}