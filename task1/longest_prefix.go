package task1

func LongestCommonPrefix(strs []string) string {
	init := false
	var pre []byte
	for _, v := range strs {
		if !init {
			pre = []byte(v)
			init = true
		} else {
			current := []byte(v)
			size := len(pre)
			l2 := len(current)

			if size > l2 {
				size = l2
			}
			tmp := make([]byte, 0)
			for i := 0; i < size; i++ {
				if current[i] == pre[i] {
					tmp = append(tmp, current[i])
				} else {
					break
				}

			}
			pre = tmp
			if len(pre) == 0 {
				break
			}
		}
	}
	return string(pre)
}
