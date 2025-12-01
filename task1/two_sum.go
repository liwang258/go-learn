package task1

func TwoSum(nums []int, target int) []int {
	var r = make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp := nums[i] + nums[j]
			if tmp == target {
				r[0] = i
				r[1] = j
			}
		}
	}
	return r
}
