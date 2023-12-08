package cp

func ExecInt(operation string, list ...int) int {
	if operation == "min" {
		min := list[0]
		for _, item := range list {
			if item < min {
				min = item
			}
		}
		return min
	}
	if operation == "max" {
		max := list[0]
		for _, item := range list {
			if item > max {
				max = item
			}
		}
		return max
	}
	panic("operation not supported")
}

func ExecLong(operation string, list ...int64) int64 {
	if operation == "min" {
		min := list[0]
		for _, item := range list {
			if item < min {
				min = item
			}
		}
		return min
	}
	if operation == "max" {
		max := list[0]
		for _, item := range list {
			if item > max {
				max = item
			}
		}
		return max
	}
	panic("operation not supported")
}

func ExecDouble(operation string, list ...float64) float64 {
	if operation == "min" {
		min := list[0]
		for _, item := range list {
			if item < min {
				min = item
			}
		}
		return min
	}
	if operation == "max" {
		max := list[0]
		for _, item := range list {
			if item > max {
				max = item
			}
		}
		return max
	}
	panic("operation not supported")
}

func ExecIntList(operation string, list []int) int {
	if operation == "min" {
		min := list[0]
		for _, item := range list {
			if item < min {
				min = item
			}
		}
		return min
	}
	if operation == "max" {
		max := list[0]
		for _, item := range list {
			if item > max {
				max = item
			}
		}
		return max
	}
	panic("operation not supported")
}

func ExecLongList(operation string, list []int64) int64 {
	if operation == "min" {
		min := list[0]
		for _, item := range list {
			if item < min {
				min = item
			}
		}
		return min
	}
	if operation == "max" {
		max := list[0]
		for _, item := range list {
			if item > max {
				max = item
			}
		}
		return max
	}
	panic("operation not supported")
}

func ExecDoubleList(operation string, list []float64) float64 {
	if operation == "min" {
		min := list[0]
		for _, item := range list {
			if item < min {
				min = item
			}
		}
		return min
	}
	if operation == "max" {
		max := list[0]
		for _, item := range list {
			if item > max {
				max = item
			}
		}
		return max
	}
	panic("operation not supported")
}

func MinInt(list ...int) int {
	min := list[0]
	for _, item := range list {
		if item < min {
			min = item
		}
	}
	return min
}

func MinLong(list ...int64) int64 {
	min := list[0]
	for _, item := range list {
		if item < min {
			min = item
		}
	}
	return min
}

func MinDouble(list ...float64) float64 {
	min := list[0]
	for _, item := range list {
		if item < min {
			min = item
		}
	}
	return min
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

func MaxInt(list ...int) int {
	max := list[0]
	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
}

func MaxLong(list ...int64) int64 {
	max := list[0]
	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
}

func MaxDouble(list ...float64) float64 {
	max := list[0]
	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
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

func GCDInt(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCMInt(a int, b int) int {
	return a / GCDInt(a, b) * b
}

func GCDLong(a int64, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCMLong(a int64, b int64) int64 {
	return a / GCDLong(a, b) * b
}
