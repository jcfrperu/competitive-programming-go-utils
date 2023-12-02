package cp

import (
	"strconv"
	"strings"
)

func Trim(s string) string {
	return strings.TrimSpace(s)
}

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

func FormatInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatLong(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FormatDouble(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
