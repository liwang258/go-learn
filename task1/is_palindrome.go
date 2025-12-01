package task1

import (
	"strconv"
)

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	remain := x
	left := 0
	length := len(strconv.Itoa(x))
	for i := 0; i < length; i++ {
		left = left*10 + remain%10
		remain = remain / 10
	}
	return left == x
}
