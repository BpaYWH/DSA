package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func next(scn *bufio.Scanner) string {
	scn.Scan()
	return scn.Text()
}

func nextInt(scn *bufio.Scanner) int {
	i, _ := strconv.Atoi(next(scn))
	return i
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	s := next(scn)

	// solve here
	res := solve(s)

	fmt.Println(res)
}

func solve(s string) string {
	runes := []rune(s)
	diff := rune('a') - rune('A')

	if runes[0] > 'Z' {
		runes[0] = runes[0] - diff
	}

	return string(runes)
}
