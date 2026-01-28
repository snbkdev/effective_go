package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func (p *PoliteChild) Eat(ctx context.Context, wg *sync.WaitGroup, chopstick1, chopstick2 *chopstick) {
	defer wg.Done()

	for {
		select{
		case <- ctx.Done():
			fmt.Printf("%s: gave up\n", p.name)
			return

		default:
			//
		}

		p.pickup(chopstick1)

		if p.sibling.IsHungry(){
			p.putDown(chopstick1)

			<- time.After(100 * time.Millisecond)
			continue
		}

		p.pickup(chopstick2)

		p.eatUntilFull()

		chopstick1.putDown()
		chopstick2.putDown()
		return
	}
}

type PoliteChild struct {
	name    string
	sibling *PoliteChild
	full    bool
}

func (p *PoliteChild) eatUntilFull() {
	p.full = true
}

func (p *PoliteChild) IsHungry() bool {
	return !p.full
}

func (p *PoliteChild) pickup(chopstick *chopstick) {
	fmt.Printf("%s: pick up\n", p.name)

	chopstick.pickup()
}

func (p *PoliteChild) putDown(chopstick *chopstick) {
	fmt.Printf("%s: put down\n", p.name)

	chopstick.putDown()
}

func NewChopstick() *chopstick {
	return &chopstick{
		mutex: &sync.Mutex{},
	}
}

type chopstick struct {
	mutex *sync.Mutex
}

func (c *chopstick) pickup() {
	c.mutex.Lock()
}

func (c *chopstick) putDown() {
	c.mutex.Unlock()
}