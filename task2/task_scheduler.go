package task2

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Delay uint //延时时间秒
	Id    int
}

func TaskSchedule() {
	record := make(map[uint]int, 0)
	lock := sync.RWMutex{}
	group := sync.WaitGroup{}
	group.Add(8)
	for i := 2; i < 10; i++ {
		task := Task{
			Delay: uint(i),
			Id:    i,
		}
		record[uint(task.Id)] = 0
		go handleTask(&group, task, &record, &lock)
	}

	group.Wait()
	for k, v := range record {
		fmt.Printf("task:%d ,expect:%d second,act:%d second \n", k, k, v)
	}

}

func handleTask(group *sync.WaitGroup, task Task, record *map[uint]int, lock *sync.RWMutex) {
	start := time.Now()

	time.AfterFunc(time.Second*time.Duration(task.Delay), func() {
		cost := int(time.Since(start).Seconds())
		lock.Lock()

		(*record)[uint(task.Id)] = cost

		lock.Unlock()

		group.Done()
	})

}
