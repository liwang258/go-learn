package main

import (
	"fmt"

	"go-learn/task1"
)

func main() {
	// nums := []int{1, 1, 5, 6, 5}
	// r := task1.SingleNumber(nums)
	// fmt.Println(r)
	// num := 5
	// v := task1.IsPalindrome(num)
	// fmt.Printf("IsPalindrome input %d,result:%v \n", num, v)

	// num = 10
	// v = task1.IsPalindrome(num)
	// fmt.Printf("IsPalindrome input %d,expect:%v, result:%v \n", num, false, v)

	// num = 12
	// v = task1.IsPalindrome(num)
	// fmt.Printf("IsPalindrome input %d,expect:%v, result:%v \n", num, false, v)
	// num = 11
	// v = task1.IsPalindrome(num)
	// fmt.Printf("IsPalindrome input %d,expect:%v, result:%v \n", num, true, v)

	// num = 121
	// v = task1.IsPalindrome(num)
	// fmt.Printf("IsPalindrome input %d,expect:%v, result:%v \n", num, true, v)
	// num = 121121
	// v = task1.IsPalindrome(num)
	// fmt.Printf("IsPalindrome input %d,expect:%v, result:%v \n", num, true, v)

	input := "()"
	v := task1.IsValid(input)
	fmt.Printf("IsValid input %s,expect:%v,result:%v \n", input, true, v)

	input = "(]"
	v = task1.IsValid(input)
	fmt.Printf("IsValid input %s,expect:%v,result:%v \n", input, false, v)

	input = ")("
	v = task1.IsValid(input)
	fmt.Printf("IsValid input %s,expect:%v,result:%v \n", input, false, v)

	input = "()[](({}))"
	v = task1.IsValid(input)
	fmt.Printf("IsValid input %s,expect:%v,result:%v \n", input, true, v)

}
