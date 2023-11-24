package cp

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func RunHackerRank(funcToRun func([]string, *bufio.Writer)) {
	run(funcToRun, os.Stdin, os.Getenv("OUTPUT_PATH")) // config for hackerrank.com
}

func RunWithString(funcToRun func([]string, *bufio.Writer), stringInput string) {
	run(funcToRun, strings.NewReader(stringInput), "")
}

func RunWithFile(funcToRun func([]string, *bufio.Writer), filePath string) {
	run(funcToRun, openFile(filePath), "")
}

func run(solutionFunc func([]string, *bufio.Writer), inputReader io.Reader, outFilePath string) {
	file, writer := getWriter(outFilePath)
	defer closeFile(file)
	defer flushWriter(writer)
	lines := getLines(inputReader)
	solutionFunc(lines, writer)
}
