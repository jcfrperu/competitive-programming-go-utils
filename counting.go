package cp

import (
	"sort"
)

func Frequencies[T int | int64 | float64 | string](list []T) (map[T]int, []T, int, int) {
	// creating map of frequencies and find indexes for max and min values
	minIndex, maxIndex := 0, 0
	frequencies := make(map[T]int, len(list))
	for i := range list {
		frequencies[list[i]]++
		if list[i] < list[minIndex] {
			minIndex = i
		}
		if list[i] > list[maxIndex] {
			maxIndex = i
		}
	}
	// list of items without repetition
	items := make([]T, 0, len(frequencies))
	for key := range frequencies {
		items = append(items, key)
	}
	// sort list of items by frequency (highest to lowest)
	sort.Slice(items, func(i, j int) bool {
		return frequencies[items[i]] > frequencies[items[j]]
	})
	return frequencies, items, minIndex, maxIndex
}
