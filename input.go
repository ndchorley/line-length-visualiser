package main

import (
	"os"
	"strings"
)

func readLines(fileName string) []string {
	return linesIn(fileName)
}

func linesIn(fileName string) []string {
	contents := readFile(fileName)

	lines := strings.Split(contents, "\n")

	return lines[0 : len(lines)-1]
}

func readFile(fileName string) string {
	bytes, _ := os.ReadFile(fileName)
	contents := string(bytes)

	return contents
}
