package cp

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	CheckError(err)
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	CheckError(err)
}

func getLines(input io.Reader) []string {
	reader := bufio.NewReaderSize(input, 16*1024*1024)
	lines := make([]string, 0, 128)

	str, _, err := reader.ReadLine()
	for err != io.EOF {
		lines = append(lines, strings.TrimRight(string(str), "\r\n"))
		str, _, err = reader.ReadLine()
	}
	return lines
}

func getWriter(filePath string) (*os.File, *bufio.Writer) {
	if len(filePath) > 0 {
		file, err := os.Create(filePath)
		CheckError(err)
		return file, bufio.NewWriterSize(file, 16*1024*1024)
	}
	return os.Stdout, bufio.NewWriterSize(os.Stdout, 16*1024*1024)
}

func flushWriter(writer *bufio.Writer) {
	err := writer.Flush()
	CheckError(err)
}

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	CheckError(err)
	return string(data)
}
