package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]

	lines := readLines(fileName)
	fmt.Println(lines)
}

func readLines(fileName string) []string {
	contents := readFile(fileName)

	lines := strings.Split(contents, "\n")

	return lines[0 : len(lines)-1]
}

func readFile(fileName string) string {
	bytes, _ := os.ReadFile(fileName)
	contents := string(bytes)

	return contents
}
