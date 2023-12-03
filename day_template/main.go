package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input_test.txt
var input string

func main() {
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		fmt.Println(s)
	}
}
