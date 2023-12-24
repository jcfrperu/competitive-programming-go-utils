package cp

import (
	"strconv"
	"strings"
)

func ParseInt(s string) int {
	value, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
	CheckError(err)
	return int(value)
}

func ParseLong(s string) int64 {
	value, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	CheckError(err)
	return value
}

func ParseDouble(s string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	CheckError(err)
	return value
}

func FormatInt[T int | int64](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatFloat[T float64](f T) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}
