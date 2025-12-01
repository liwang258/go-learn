package task1

// 删除有序数组中的重复项目
// 见 leetcode 26
func RemoveDuplicates(nums []int) int {
	record := make(map[int]bool, 0)
	size := len(nums)
	// var b []byte
	for i := 0; i < size; i++ {
		_, exist := record[nums[i]]
		if exist {
			for idx := i + 1; idx < len(nums); idx++ {
				nums[idx-1] = nums[idx]
			}
			//尾部重复的就不要去遍历了
			size = size - 1
			//因为元素全部从后往前挪了一位，因此需要继续从当前位置扫一遍
			i = i - 1
			//b, _ = json.Marshal(nums)
			//fmt.Printf("i:%d v:%d nums:%s \n", i, nums[i], string(b))
		} else {
			record[nums[i]] = true
		}
	}
	// b, _ = json.Marshal(nums)
	// fmt.Printf("parsed:%s \n", string(b))
	return len(record)
}
