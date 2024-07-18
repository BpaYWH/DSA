package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	k := nextInt(scn)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt(scn)
	}

	// solve here
	solve(n, k, x)
}

func solve(n int, k int, x []int) {
	if k == 0 {
		if x[0] > 1 {
			fmt.Println(1)
		} else {
			fmt.Println(-1)
		}
		return
	}

	if k == n {
		fmt.Println(x[k - 1])
		return
	}

	sort.Ints(x)

	if x[k - 1] == x[k] {
		fmt.Println(-1)
	} else {
		fmt.Println(x[k - 1])
	}
}