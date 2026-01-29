package main

func Example() bool {
	doneCh := make(chan struct{}, 10)

	for x := 0; x < 10; x++ {
		go func ()  {
			defer func() {
				doneCh <- struct{}{}
			}()

			//
		}()
	}

	for x := 0; x < 10; x++ {
		<- doneCh
	}

	return true
}