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

func ReadLine(path string) string {
  lines := ReadLines(path)
  if len(lines) != 1 {
    panic("not a single line file")
  }
  return lines[0]
}
