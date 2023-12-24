package cp

import "cmp"

func Find[T cmp.Ordered](list []T, value T) int {
	for i := range list {
		if list[i] == value {
			return i
		}
	}
	return -1
}

func FindAll[T cmp.Ordered](list []T, value T) []int {
	matches := make([]int, 0)
	for i := range list {
		if list[i] == value {
			matches = append(matches, i)
		}
	}
	return matches
}

func FindLast[T cmp.Ordered](list []T, value T) int {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == value {
			return i
		}
	}
	return -1
}
