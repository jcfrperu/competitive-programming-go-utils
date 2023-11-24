package cp

import (
	"sort"
)

func FrequencesInt(list []int) ([]int, map[int]int) {
	// creating map of frequencies
	freqMap := make(map[int]int, len(list))
	for _, item := range list {
		freqMap[item] += 1
	}
	// sorted list of items without repetition
	sortedItems := make([]int, 0, len(freqMap))
	for key, _ := range freqMap {
		sortedItems = append(sortedItems, key)
	}
	sort.Ints(sortedItems)
	return sortedItems, freqMap
}

func FrequencesString(list []string) ([]string, map[string]int) {
	// creating map of frequencies
	freqMap := make(map[string]int, len(list))
	for _, item := range list {
		freqMap[item] += 1
	}
	// sorted list of items without repetition
	sortedItems := make([]string, 0, len(freqMap))
	for key, _ := range freqMap {
		sortedItems = append(sortedItems, key)
	}
	sort.Strings(sortedItems)
	return sortedItems, freqMap
}
