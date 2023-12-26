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

func Min[T int | int64 | float64 | string](list []T) T {
	return list[MinIndex(list)]
}

func MinIndex[T int | int64 | float64 | string](list []T) int {
	minIndex := 0
	for i := range list {
		if list[i] < list[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func Max[T int | int64 | float64 | string](list []T) T {
	return list[MaxIndex(list)]
}

func MaxIndex[T int | int64 | float64 | string](list []T) int {
	maxIndex := 0
	for i := range list {
		if list[i] > list[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func MinMax[T int | int64 | float64 | string](list []T) (T, T) {
	minIndex, maxIndex := MinMaxIndex(list)
	return list[minIndex], list[maxIndex]
}

func MinMaxIndex[T int | int64 | float64 | string](list []T) (int, int) {
	minIndex := 0
	maxIndex := 0
	for i := range list {
		if list[i] < list[minIndex] {
			minIndex = i
		}
		if list[i] > list[maxIndex] {
			maxIndex = i
		}
	}
	return minIndex, maxIndex
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
