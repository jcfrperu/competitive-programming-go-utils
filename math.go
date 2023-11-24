package cp

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxLong(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MaxDouble(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func MaxIntList(list []int) int {
	max := list[0]
	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
}

func MaxLongList(list []int64) int64 {
	max := list[0]
	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
}

func MaxDoubleList(list []float64) float64 {
	max := list[0]
	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinLong(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MinIntList(list []int) int {
	min := list[0]
	for _, item := range list {
		if item < min {
			min = item
		}
	}
	return min
}

func MinLongList(list []int64) int64 {
	min := list[0]
	for _, item := range list {
		if item < min {
			min = item
		}
	}
	return min
}

func MinDoubleList(list []float64) float64 {
	min := list[0]
	for _, item := range list {
		if item < min {
			min = item
		}
	}
	return min
}
