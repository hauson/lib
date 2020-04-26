package static

/*
math statistic
 */

func MinInts(datas ...int) int {
	if len(datas) == 0 {
		panic("slice is nil")
	}

	min := datas[0]
	for _, num := range datas[1:] {
		if num < min {
			min = num
		}
	}

	return min
}

func MaxInts(datas ...int) int {
	if len(datas) == 0 {
		panic("slice is nil")
	}
	max := datas[0]
	for _, num := range datas[1:] {
		if num > max {
			max = num
		}
	}

	return max
}
