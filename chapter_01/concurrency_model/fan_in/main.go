package main

func FanInExample(stopCh chan struct{}, inputChA, inputChB chan int, outputCh chan int) {
	for{
		var data int
		var isOpen bool

		select{
		case data, isOpen = <- inputChA:
			if !isOpen {
				inputChA = nil
			}
		case data, isOpen = <- inputChB:
			if !isOpen {
				inputChB = nil
			}
		case <- stopCh:
			return
		}

		if !isOpen {
			if inputChA == nil && inputChB == nil {
				return
			}

			continue
		}

		outputCh <- data
	}
}