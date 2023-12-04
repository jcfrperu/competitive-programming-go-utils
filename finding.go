package cp

func FindInList(list []string, value string) int {
	for i, item := range list {
		if value == item {
			return i
		}
	}
	return -1
}

func FindInListInt(list []int, value int) int {
	for i, item := range list {
		if value == item {
			return i
		}
	}
	return -1
}

func FindInListLong(list []int64, value int64) int {
	for i, item := range list {
		if value == item {
			return i
		}
	}
	return -1
}

func FindInListDouble(list []float64, value float64) int {
	for i, item := range list {
		if value == item {
			return i
		}
	}
	return -1
}
