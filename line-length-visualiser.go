package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	contents := readFile(fileName)
	fmt.Println(contents)
}

func readFile(fileName string) string {
	bytes, _ := os.ReadFile(fileName)
	contents := string(bytes)

	return contents
}
