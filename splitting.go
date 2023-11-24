package cp

import "strings"

func Split(line string, separator string) []string {
	return strings.Split(line, separator)
}

func SplitInt(line string, separator string) []int {
	items := Split(line, separator)
	list := make([]int, 0, len(items))
	for _, item := range items {
		list = append(list, ParseInt(item))
	}
	return list
}

func SplitLong(line string, separator string) []int64 {
	items := Split(line, separator)
	list := make([]int64, 0, len(items))
	for _, item := range items {
		list = append(list, ParseLong(item))
	}
	return list
}

func SplitDouble(line string, separator string) []float64 {
	items := Split(line, separator)
	list := make([]float64, 0, len(items))
	for _, item := range items {
		list = append(list, ParseDouble(item))
	}
	return list
}

func GetInt(line string, part int) int {
	split := Split(line, " ")
	return ParseInt(split[part])
}

func GetLong(line string, part int) int64 {
	split := Split(line, " ")
	return ParseLong(split[part])
}

func GetDouble(line string, part int) float64 {
	split := Split(line, " ")
	return ParseDouble(split[part])
}
