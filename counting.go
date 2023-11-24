package cp

import (
	"sort"
)

func FrequenciesInt(list []int) ([]int, map[int]int) {
	// creating map of frequencies
	freqMap := make(map[int]int, len(list))
	for _, item := range list {
		freqMap[item] += 1
	}
	// sorted list of items without repetition
	sortedItems := make([]int, 0, len(freqMap))
	for key := range freqMap {
		sortedItems = append(sortedItems, key)
	}
	sort.Ints(sortedItems)
	return sortedItems, freqMap
}

func FrequenciesString(list []string) ([]string, map[string]int) {
	// creating map of frequencies
	freqMap := make(map[string]int, len(list))
	for _, item := range list {
		freqMap[item] += 1
	}
	// sorted list of items without repetition
	sortedItems := make([]string, 0, len(freqMap))
	for key := range freqMap {
		sortedItems = append(sortedItems, key)
	}
	sort.Strings(sortedItems)
	return sortedItems, freqMap
}
