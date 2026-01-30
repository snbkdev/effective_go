package main

import (
	"fmt"
	"sync"
)

func Example() {
	inputs := []int{1, 2, 3, 4, 5, 6}

	outputs := make([]int, len(inputs))

	wg := &sync.WaitGroup{}
	wg.Add(len(inputs))

	for index, value := range inputs {
		go func ()  {
			defer wg.Done()

			outputs[index] = value * 2
		}()
	}

	wg.Wait()

	for index, value := range outputs {
		fmt.Printf("%d -> %d\n", index, value)
	}
}

func Example2() {
	inputs := []int{1, 2, 3, 4, 5, 6, 7}

	outputs := make([]int, len(inputs))
	outputMutex := &sync.Mutex{}

	wg := &sync.WaitGroup{}
	wg.Add(len(inputs))

	for index, value := range inputs {
		go func()  {
			defer wg.Done()

			outputMutex.Lock()
			outputs[index] = value * 2
			outputMutex.Unlock()
		}()
	}

	wg.Wait()

	for index, value := range outputs {
		fmt.Printf("%d -> %d\n", index, value)
	}
}

func main() {
	//Example()
}