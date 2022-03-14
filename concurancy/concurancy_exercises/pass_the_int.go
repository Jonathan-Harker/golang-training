package concurancy_exercises

import "sync"

func PassTheInt() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	routines := 2
	wg.Add(routines)

	for i := 1; i <= routines; i++ {
		runRoutine(ch, &wg)
	}

	ch <- 1
	wg.Wait()
}

func runRoutine(ch chan int, wg *sync.WaitGroup) {
	go func() {
		goroutine(ch)
		wg.Done()
	}()
}

func goroutine(ch chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			return
		}

		println(i)

		if i == 10 {
			close(ch)
			return
		}

		i++
		ch <- i
	}
}
