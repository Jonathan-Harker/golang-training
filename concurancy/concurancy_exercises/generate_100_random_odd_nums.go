package concurancy_exercises

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func Generate100RandomOddNums() {
	values := make(chan int)
	closeChannel := make(chan bool, 1)
	poolSize := runtime.GOMAXPROCS(0)
	wg := sync.WaitGroup{}
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		values = sendInts(values, closeChannel, &wg, i)
	}

	numbers := receiveInts(values)

	close(closeChannel)
	wg.Wait()
	fmt.Println(numbers)
}

func sendInts(values chan int, closeChannel chan bool, wg *sync.WaitGroup, i int) chan int {
	go func(id int) {
		for {
			num := rand.Intn(1000)
			select {
			case values <- num:
				fmt.Println("")
			case <-closeChannel:
				wg.Done()
				return
			}
		}
	}(i)
	return values
}

func receiveInts(values chan int) []int {
	var numbers []int

	for value := range values {
		if value%2 == 0 {
			continue
		}

		numbers = append(numbers, value)
		if len(numbers) == 100 {
			break
		}
	}
	return numbers
}
