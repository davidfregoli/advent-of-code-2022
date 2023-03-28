package reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) []string {
	var lines []string
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()
	return lines
}
