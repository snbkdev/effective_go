package main

import (
	"context"

	"golang.org/x/sync/semaphore"
)

const (
	lightWorkCost = 1
	heavyWorkCost = 4

	maxConcurrency = 10
)

func Example(stopCh chan struct{}, lightWorkCh chan Data, heavyWorkCh chan Data) {
	weightedSemaphore := semaphore.NewWeighted(maxConcurrency)

	for {
		select {
		case data := <- lightWorkCh:
			go func() {
				err := weightedSemaphore.Acquire(context.Background(), lightWorkCost)

				if err != nil {
					return
				}

				doLightWork(data)

				weightedSemaphore.Release(lightWorkCost)
			}()

			case data := <- heavyWorkCh:
				go func() {
					err := weightedSemaphore.Acquire(context.Background(), heavyWorkCost)

					if err != nil {
						return
					}

					doHeavyWork(data)

					weightedSemaphore.Release(heavyWorkCost)
				}()

				case <- stopCh:
					_ = weightedSemaphore.Acquire(context.Background(), maxConcurrency)
					return
		}
	}
}

func doLightWork(data Data) {}

func doHeavyWork(data Data) {}

type Data struct {}