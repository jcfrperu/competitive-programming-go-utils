package cp

import (
	"strconv"
	"strings"
)

func IsInt(s string) bool {
	_, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
	return err == nil
}

func IsLong(s string) bool {
	_, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return err == nil
}

func IsDouble(s string) bool {
	_, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return err == nil
}
