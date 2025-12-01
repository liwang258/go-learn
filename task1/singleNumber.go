package task1

func SingleNumber(nums []int) int {
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
