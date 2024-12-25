package cp

import "sort"

func Freq[T int | int64 | float64 | string](list []T, sortByFreq bool) (map[T]int, []T, int, int) {
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

	if sortByFreq {
		// sort list of items by frequency (highest to lowest)
		sort.Slice(items, func(i, j int) bool {
			return frequencies[items[i]] > frequencies[items[j]]
		})
	} else {
		// natural sorting (lowest to highest)
		sort.Slice(items, func(i, j int) bool {
			return items[i] < items[j]
		})
	}
	return frequencies, items, minIndex, maxIndex
}

func Permute[T any](items []T, k int, createCopy bool) [][]T {
	initCapacity := Pow(len(items), k) // be careful with big values of 'k' and len(items)
	permutations := make([][]T, 0, initCapacity)

	// helper function to generate permutations in groups of k elements
	var helper func([]T, int)
	helper = func(current []T, k int) {
		if k == 0 {
			permutation := current
			if createCopy {
				permutation = make([]T, len(current))
				copy(permutation, current)
			}
			permutations = append(permutations, permutation)
			return
		}
		for i := 0; i < len(items); i++ {
			helper(append(current, items[i]), k-1)
		}
	}

	if k < 0 {
		// all permutations of sizes 1, 2, 3, ... len(items) in one big list
		for i := 1; i <= len(items); i++ {
			helper([]T{}, i)
		}
	} else {
		// all permutations of k 'k'
		helper([]T{}, k)
	}
	return permutations
}
