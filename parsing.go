package cp

import "strconv"

func ParseInt(s string) int {
	value, err := strconv.ParseInt(s, 10, 32)
	CheckError(err)
	return int(value)
}

func ParseLong(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	CheckError(err)
	return value
}

func ParseDouble(s string) float64 {
	value, err := strconv.ParseFloat(s, 64)
	CheckError(err)
	return value
}
