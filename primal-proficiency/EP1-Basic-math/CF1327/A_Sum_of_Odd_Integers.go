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
		n := nextInt(scn)
		k := nextInt(scn)

		solve(n, k, 0)
	}
}

func solve(n, k, i int) {
	if (n % 2) != (k % 2) {
		fmt.Println("NO")

		return
	}

	if k == 1 {
		if (2 * i + 1) <= n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		return
	}

	if k > 1 {
		solve(n - (2 * i + 1), k - 1, i + 1)
	}
}