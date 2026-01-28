package main

func FanOutExample() {
	inputCh := make(chan int)
	outputChA := make(chan int)
	outputChB := make(chan int)

	for data := range inputCh {
		select {
		case outputChA <- data:
			
		case outputChB <- data:
		}
	}
}