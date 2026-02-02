package main

import (
	"context"
	"net"
	"runtime"
)

func NewPool(network, address string, max int) *Pool {
	return &Pool{
		network: network,
		address: address,

		connections: make(chan net.Conn, max),
		semaphore: make(chan struct{}, max),
	}
}

type Pool struct {
	network string
	address string

	connections chan net.Conn
	semaphore chan struct{}
}

func (p *Pool) Get(ctx context.Context) (net.Conn, error) {
	for {
		select {
		case conn := <- p.connections:
			return conn, nil
		default:
		}

		select {
		case p.semaphore <- struct{}{}:
			conn, err := net.Dial(p.network, p.address)
			if err != nil {
				<- p.semaphore

				return nil, err
			}

			return conn, nil

		case <- ctx.Done():
			return nil, ctx.Err()

		default:
			runtime.Gosched()
		}
	}
}

func (p *Pool) Release(conn net.Conn) {
	p.connections <- conn
}