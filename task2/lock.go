package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Lock() {
	group := sync.WaitGroup{}
	group.Add(10)
	l := sync.Mutex{}
	val := 0
	for i := 0; i < 10; i++ {
		go add(&group, &l, &val)
	}
	group.Wait()
	fmt.Printf("val:%d \n", val)
}

func NoLock() {
	group := sync.WaitGroup{}
	group.Add(10)
	var val int32 = 0
	for i := 0; i < 10; i++ {
		go addNoLock(&group, &val)
	}
	group.Wait()
	fmt.Printf("val:%d \n", val)
}

func addNoLock(group *sync.WaitGroup, val *int32) {
	defer group.Done()
	for i := 1; i <= 1000; i++ {
		atomic.AddInt32(val, 1)
	}

}

func add(group *sync.WaitGroup, lock *sync.Mutex, val *int) {
	defer group.Done()
	lock.Lock()
	for i := 1; i <= 1000; i++ {
		*val = *val + 1
	}
	defer lock.Unlock()

}
