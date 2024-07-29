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

	numCases := nextInt(scn)

	for i := 0; i < numCases; i++ {
		l := nextInt(scn)
		r := nextInt(scn)

		solve(l, r)
	}
}

func solve(l, r int) {
	if 2 * l > r {
		fmt.Println(-1, -1)
	} else {
		fmt.Println(l, 2 * l)
	}
}