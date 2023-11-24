package cp

import (
	"bufio"
	"fmt"
	"strings"
)

func Out(writer *bufio.Writer, data string) {
	PrintOut(writer, data, true, true)
}

func PrintOut(writer *bufio.Writer, data string, trimLeft bool, trimRight bool) {
	dataToPrint := data
	if trimLeft {
		dataToPrint = strings.TrimLeft(dataToPrint, " ")
	}
	if trimRight {
		dataToPrint = strings.TrimRight(dataToPrint, " ")
	}
	_, err := fmt.Fprintf(writer, dataToPrint)
	CheckError(err)
}
