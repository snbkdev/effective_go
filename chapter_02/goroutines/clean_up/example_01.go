package main

func Example01() {
	go func() {
		for {
			doSomething()
	}
	}()
}

func doSomething() {}