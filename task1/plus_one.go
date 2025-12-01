package task1

// https://leetcode.cn/problems/plus-one/
// see leetcode =66
func PlusOne(digits []int) []int {
	addtion := 0
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i]
		if i == len(digits)-1 {
			if v < 9 {
				digits[i] = digits[i] + 1
				addtion = 0
			} else {
				digits[i] = 0
				addtion = 1
			}
		}

		if (v + addtion) <= 9 {
			digits[i] = digits[i] + addtion
			addtion = 0
			break
		} else {
			digits[i] = 0
			addtion = 1
		}
	}
	var result []int
	if addtion == 1 {
		result = append(result, 1)
	}
	result = append(result, digits...)

	return result
}
