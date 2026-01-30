package main

import (
	"errors"
	"time"
)

func readWithTimeout(closure func() string, timeout time.Duration) (string, error) {
	resultCh := make(chan string, 1)
	go func ()  {
		resultCh <- closure()
	}()

	select {
	case result := <- resultCh:
		
		return result, nil

	case <- time.After(timeout):
		return "", errors.New("read timed out")
	}
}