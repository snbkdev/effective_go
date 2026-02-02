package main

type Clock struct {
	listeners []chan int64
}

func (c *Clock) AddListener(in chan int64) {
	c.listeners = append(c.listeners, in)
}

func (c *Clock) onTick(ts int64) {
	for _, listener := range c.listeners {
		select {
		case listener <- ts:
			
		default:
			
		}
	}
}