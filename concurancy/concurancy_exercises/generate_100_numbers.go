package concurancy_exercises

import (
	"fmt"
	"math/rand"
	"time"
)

const numRoutines = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Generate100Numbers() {
	ch := make(chan int, numRoutines)
	sendNumbers(ch)
	msgs := receiveNumbers(ch)
	fmt.Println(msgs)
}

func sendNumbers(ch chan int) {
	for i := 0; i < numRoutines; i++ {
		go func() {
			ch <- rand.Intn(1000)
		}()
	}
}

func receiveNumbers(ch chan int) []int {
	var msgs []int

	for i := 0; i < numRoutines; i++ {
		msg, _ := <-ch
		msgs = append(msgs, msg)
	}

	return msgs
}
