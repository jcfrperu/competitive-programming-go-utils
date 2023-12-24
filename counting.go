package cp

import (
	"cmp"
	"sort"
)

func Frequencies[T cmp.Ordered](list []T, sortKeys bool) (map[T]int, []T) {
	// creating map of frequencies
	freqMap := make(map[T]int, len(list))
	for i := range list {
		freqMap[list[i]]++
	}
	// sorted list of items without repetition
	keys := make([]T, 0, len(freqMap))
	for key := range freqMap {
		keys = append(keys, key)
	}
	if sortKeys {
		sort.Slice(keys, func(i, j int) bool {
			return freqMap[keys[i]] > freqMap[keys[j]]
		})
	}
	return freqMap, keys
}
