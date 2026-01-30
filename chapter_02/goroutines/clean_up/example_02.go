package main

func Example02() {
	stopCh := make(chan struct{})
	defer close(stopCh)

	go func() {
		for {
			doSomething()

			select{
			case <- stopCh:
				return
			default:
				//
			}
		}
	}()
}