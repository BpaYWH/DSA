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

	n := nextInt(scn)

	// solve here
	solve(n)

	return
}

func solve(n int) {
	numStr := strconv.Itoa(n)
	numRunes := []rune(numStr)

	for i := range numRunes {
		if (i == 0 && numRunes[i] == '9') || numRunes[i] < '5' {
			continue
		}

		numRunes[i] = rune(int('9') - int(numRunes[i]) + int('0'))
	}

	numStr = string(numRunes)

	fmt.Println(numStr)
}
