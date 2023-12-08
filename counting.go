package cp

import (
	"sort"
)

func Frequencies(list []string, sortKeys bool) (map[string]int, []string) {
	// creating map of frequencies
	freqMap := make(map[string]int, len(list))
	for _, item := range list {
		freqMap[item]++
	}
	// sorted list of items without repetition
	keys := make([]string, 0, len(freqMap))
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

func FrequenciesInt(list []int, sortKeys bool) (map[int]int, []int) {
	// creating map of frequencies
	freqMap := make(map[int]int, len(list))
	for _, item := range list {
		freqMap[item]++
	}
	// sorted list of items without repetition
	keys := make([]int, 0, len(freqMap))
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

func FrequenciesLong(list []int64, sortKeys bool) (map[int64]int, []int64) {
	// creating map of frequencies
	freqMap := make(map[int64]int, len(list))
	for _, item := range list {
		freqMap[item]++
	}
	// sorted list of items without repetition
	keys := make([]int64, 0, len(freqMap))
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
