package main

func ExampleReadAllNoBoolean() {
	events := make(chan Event)

	for {
		event := <- events

		processEvent(event)
	}
}

func ExampleReadAllWithBoolean() {
	events := make(chan Event)

	for {
		event, isClosed := <- events
		if isClosed {
			return
		}

		processEvent(event)
	}
}

func ExampleReadAllRange() {
	events := make(chan Event)

	for event := range events {
		processEvent(event)
	}
}

func processEvent(even Event) {}

type Event struct {}