package main

import "time"

func main() {
	table := make(chan *ball)

	go newPlayer("paul: ping", table)
	go newPlayer("sally: pong", table)

	table <- &ball{}

	<- time.After(3 * time.Second)
}

func newPlayer(action string, table chan *ball) {
	for thisBall := range table{
		println(action)

		table <- thisBall
	}
}

type ball struct {
	hits int
}