package cp

func GCD[T int | int64](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM[T int | int64](a, b T) T {
	return a / GCD(a, b) * b
}

func Min[T int | int64 | float64 | string](list []T) (T, int) {
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

func Max[T int | int64 | float64 | string](list []T) (T, int) {
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

func MinMax[T int | int64 | float64 | string](list []T) (T, int, T, int) {
	minIndex, minValue := 0, list[0]
	maxIndex, maxValue := 0, list[0]
	for i := range list {
		if list[i] < minValue {
			minIndex = i
			minValue = list[i]
		}
		if list[i] > maxValue {
			maxIndex = i
			maxValue = list[i]
		}
	}
	return minValue, minIndex, maxValue, maxIndex
}

func AsList[T int | int64 | float64 | string](a ...T) []T {
	list := make([]T, 0, len(a))
	if len(a) > 0 {
		for i := range a {
			list = append(list, a[i])
		}
	}
	return list
}
