package cp

func GCD[T IntegerNumber](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM[T IntegerNumber](a, b T) T {
	return a / GCD(a, b) * b
}

func MinInt(list []int) (int, int) {
	index := 0
	minValue := list[0]
	for i := range list {
		if list[i] < minValue {
			index = i
			minValue = list[i]
		}
	}
	return minValue, index
}

func MinLong(list []int64) (int64, int) {
	index := 0
	minValue := list[0]
	for i := range list {
		if list[i] < minValue {
			index = i
			minValue = list[i]
		}
	}
	return minValue, index
}

func MinDouble(list []float64) (float64, int) {
	index := 0
	minValue := list[0]
	for i := range list {
		if list[i] < minValue {
			index = i
			minValue = list[i]
		}
	}
	return minValue, index
}

func MaxInt(list []int) (int, int) {
	index := 0
	maxValue := list[0]
	for i := range list {
		if list[i] > maxValue {
			index = i
			maxValue = list[i]
		}
	}
	return maxValue, index
}

func MaxLong(list []int64) (int64, int) {
	index := 0
	maxValue := list[0]
	for i := range list {
		if list[i] > maxValue {
			index = i
			maxValue = list[i]
		}
	}
	return maxValue, index
}

func MaxDouble(list []float64) (float64, int) {
	index := 0
	maxValue := list[0]
	for i := range list {
		if list[i] > maxValue {
			index = i
			maxValue = list[i]
		}
	}
	return maxValue, index
}

func AsListInt(a ...int) []int {
	list := make([]int, 0, len(a))
	if len(a) > 0 {
		for i := range a {
			list = append(list, a[i])
		}
	}
	return list
}

func AsListLong(a ...int64) []int64 {
	list := make([]int64, 0, len(a))
	if len(a) > 0 {
		for i := range a {
			list = append(list, a[i])
		}
	}
	return list
}

func AsListDouble(a ...float64) []float64 {
	list := make([]float64, 0, len(a))
	if len(a) > 0 {
		for i := range a {
			list = append(list, a[i])
		}
	}
	return list
}
