package task2

import (
	"fmt"
	"sync"
)

func ChannelWithNoBuffer() {
	group := sync.WaitGroup{}
	group.Add(2)
	ch := make(chan (int))
	go Send(&group, 10, &ch)
	go Read(&group, &ch)
	group.Wait()
}

func ChannelWithBuffer() {
	group := sync.WaitGroup{}
	group.Add(2)
	ch := make(chan (int), 5)
	go Send(&group, 100, &ch)
	go Read(&group, &ch)
	group.Wait()
}

func Send(group *sync.WaitGroup, limit int, ch *chan (int)) {
	for i := 1; i <= limit; i++ {
		*ch <- i
	}
	close(*ch)
	defer group.Done()
}

func Read(group *sync.WaitGroup, ch *chan (int)) {
	for v := range *ch {
		fmt.Printf("read %d from channel \n", v)
	}
	defer group.Done()
}
