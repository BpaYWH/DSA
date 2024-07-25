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
	const maxBufSize = 1024 * 1024
	buf := make([]byte, maxBufSize)
	scn := bufio.NewScanner(os.Stdin)
	scn.Buffer(buf, maxBufSize)
	scn.Split(bufio.ScanWords)

	nextInt(scn)
	inputStr := next(scn)

	// solve here
	solve(inputStr)
}

func solve(s string) {
	for i := 1; i < len(s); i++ {
		if (s[i] < s[i-1]) {
			fmt.Print(string(s[:i-1] + s[i:]))
			return
		}
	}

	fmt.Print(s[:len(s)-1])
}