package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	bytes, _ := os.ReadFile(fileName)
	contents := string(bytes)
	fmt.Println(contents)
}
