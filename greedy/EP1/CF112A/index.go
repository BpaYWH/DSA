package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input1 := strings.ToLower(scanner.Text())
	scanner.Scan()
	input2 := strings.ToLower(scanner.Text())

	if input1 == input2 {
		fmt.Println(0)
		return
	}
	if input1 > input2 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
