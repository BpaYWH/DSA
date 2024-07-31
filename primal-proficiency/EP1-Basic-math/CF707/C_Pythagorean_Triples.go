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

	n := nextInt(scn)

	solve(n)
}

func solve(n int) {
	if n <= 2 {
		fmt.Println(-1)
		return
	}

	if n % 2 == 1 {
		b := n * n / 2
		c := n * n / 2 + 1
		fmt.Print(b, c)
	} else {
		b := n * n / 4 - 1
		c := n * n / 4 + 1
		fmt.Print(b, c)
	}
}