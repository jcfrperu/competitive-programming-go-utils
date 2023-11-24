package cp

import "strings"

func Split(line string, separator string) []string {
	return strings.Split(line, separator)
}

func SplitInts(line string, separator string) []int {
	items := Split(line, separator)
	list := make([]int, 0, len(items))
	for _, item := range items {
		list = append(list, ParseInt(item))
	}
	return list
}

func SplitLongs(line string, separator string) []int64 {
	items := Split(line, separator)
	list := make([]int64, 0, len(items))
	for _, item := range items {
		list = append(list, ParseLong(item))
	}
	return list
}

func SplitDoubles(line string, separator string) []float64 {
	items := Split(line, separator)
	list := make([]float64, 0, len(items))
	for _, item := range items {
		list = append(list, ParseDouble(item))
	}
	return list
}

func GetSplitInt(line string, separator string, part int) int {
	split := Split(line, separator)
	return ParseInt(split[part])
}

func GetSplitLong(line string, separator string, part int) int64 {
	split := Split(line, separator)
	return ParseLong(split[part])
}

func GetSplitDouble(line string, separator string, part int) float64 {
	split := Split(line, separator)
	return ParseDouble(split[part])
}
