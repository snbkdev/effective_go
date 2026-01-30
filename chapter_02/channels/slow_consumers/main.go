package main

func Example(signalCh chan struct{}) {
	select {
	case signalCh <- struct{}{}:
		//send signal
		
	default:
		// drop signal as one is already pending
	}
}