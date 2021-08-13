package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("usage go run main.go 'go is awesome'")
	}
	str := os.Args[1]
	cnt := WordCount(str)
	fmt.Println(cnt)
}

func WordCount(str string) int {
	words := strings.Split(str, " ")
	return len(words)
}
