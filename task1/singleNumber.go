package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var target int
	record := make(map[int]int)
	for _, v := range nums {
		_, ok := record[v]
		if !ok {
			record[v] = 1
		} else {
			delete(record, v)
		}
	}
	for v, n := range record {
		if n == 1 {
			target = v
			break
		}
	}
	return target
}

func main() {
	nums := []int{1, 1, 5, 6, 5}
	r := singleNumber(nums)
	fmt.Println(r)
}
