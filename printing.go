package cp

import (
	"bufio"
	"fmt"
)

func Print(writer *bufio.Writer, format string, a ...any) {
	_, err := fmt.Fprintf(writer, format, a)
	CheckError(err)
}
