package cp

import (
	"bufio"
	"io"
	"os"
)

func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	CheckError(err)
	return file
}

func CloseFile(file *os.File) {
	err := file.Close()
	CheckError(err)
}

func GetLines(input io.Reader) []string {
	reader := bufio.NewReaderSize(input, 16*1024*1024)
	var lines = make([]string, 0, 128)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func GetWriter(outputPath string) (*os.File, *bufio.Writer) {
	if len(outputPath) > 0 {
		file, err := os.Create(outputPath)
		CheckError(err)
		return file, bufio.NewWriterSize(file, 16*1024*1024)
	}
	return os.Stdout, bufio.NewWriterSize(os.Stdout, 16*1024*1024)
}

func FlushWriter(writer *bufio.Writer) {
	err := writer.Flush()
	CheckError(err)
}

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	CheckError(err)
	return string(data)
}
