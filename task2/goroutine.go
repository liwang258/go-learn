package task2

import (
	"fmt"
	"sync"
)

func Goroutine_Print() {
	group := sync.WaitGroup{}
	group.Add(2)
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	go printCh1(ch1, &group)
	go printCh2(ch2, &group)

	for i := 1; i <= 10; i++ {
		ch1 <- i
		ch2 <- i
	}
	close(ch1)
	close(ch2)
	group.Wait()
}

func printCh1(ch chan (int), group *sync.WaitGroup) {
	defer group.Done()
	for v := range ch {
		if v%2 == 0 {
			fmt.Printf("ch1,print %d \n", v)
		}
	}

}

func printCh2(ch chan int, group *sync.WaitGroup) {
	defer group.Done()
	for v := range ch {
		if v%2 == 1 {
			fmt.Printf("ch2,print %d \n", v)
		}
	}
}
