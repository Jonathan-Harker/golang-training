package concurancy_exercises

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const goroutines = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomOddNums() {
	ch := make(chan int, goroutines)
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			num := rand.Intn(1000)

			if num%2 == 0 {
				return
			}

			ch <- num
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var ints []int
	for i := range ch {
		ints = append(ints, i)
	}

	fmt.Println(ints)
}
