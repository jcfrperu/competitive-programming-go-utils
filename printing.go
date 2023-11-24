package cp

import (
	"bufio"
	"fmt"
	"strings"
)

// Print prints a string using a Writer
func Print(writer *bufio.Writer, data string) {
	_, err := fmt.Fprintf(writer, data)
	CheckError(err)
}

// Out prints a right trimmed string using a writer (specific to competitive programming platforms - avoid extra spaces at the end)
func Out(writer *bufio.Writer, data string) {
	trimRight := strings.TrimRight(data, " ")
	Print(writer, trimRight)
}
