package main

import (
	"fmt"
	"strings"
)

func main() {
	var statements int = 0
	var input string = ""
	var res int = 0

	fmt.Scanln(&statements)

	for i := 0; i < statements; i++ {
		fmt.Scanln(&input)

		if strings.Contains(input, "++") {
			res++
		}
		if strings.Contains(input, "--") {
			res--
		}
	}

	fmt.Println(res)
}
