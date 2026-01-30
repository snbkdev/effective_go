package main

func readFromNilChannel() []string {
	var dataCh chan struct{}
	var results []string

	for x := 0; x < 4; x++ {
		select {
		case <- dataCh:
			results = append(results, "read from channel")

			dataCh = nil
		default:
			results = append(results, "default")

			dataCh = make(chan struct{}, 1)
			dataCh <- struct{}{}
		}
	}

	return results
}