package task1

import (
	"strings"
)

func IsValid(s string) bool {
	left := make(map[byte]byte, 6)
	left['('] = ')'
	left['{'] = '}'
	left['['] = ']'

	s = strings.TrimSpace(s)
	outPut := make([]byte, 0)
	outPut = append(outPut, s...)
	stack := make([]byte, 0)
	//fmt.Printf("target out put:%s \n", string(outPut))
	for _, v := range outPut {
		//判断是否是左边符号，如果是左边符号则需要入栈，否则执行出栈判断
		_, isInput := left[v]
		//fmt.Printf("start val:%s,is exist:%v \n", string(v), isInput)
		if isInput {
			stack = append(stack, left[v])
		} else {

			//如果栈中没有数据，则说明前面的都匹配完毕了，这个时候如果出现的是一个右边符号，则直接不匹配
			if len(stack) == 0 {
				return false
			}
			//获取栈中最后一个入栈的元素 转换成目标的出栈元素
			shouldbe := stack[len(stack)-1]
			//fmt.Printf("stack:%s,current out element:%s,last input element:%s \n", string(stack), string(v), string(lastInput))
			//目标出栈元素不和当前右符号不匹配，则直接返回false
			if shouldbe != v {
				return false
			} else {
				//能匹配，则吧上一个已经入栈的左边元素 出栈
				stack = stack[0 : len(stack)-1]
			}

		}

	}

	return len(stack) == 0
}
