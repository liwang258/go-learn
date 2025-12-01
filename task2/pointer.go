package task2

func Add(v *int) {
	*v = *v + 10
}

func SilencOpt(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}
