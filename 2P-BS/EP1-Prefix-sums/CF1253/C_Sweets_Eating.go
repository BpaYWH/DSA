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
	m := nextInt(scn)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt(scn)
	}

	// solve here
	solve(n, m, x)
}

func solve(n, m int, arr []int) {
	sort.Ints(arr)
	preA := make([]int, n + 1)
	preA[0] = 0
	for i, _ := range arr {
		preA[i + 1] = arr[i] + preA[i]
	}

	prePreA := make([]int, n + 1)
	for i, _ := range preA {
		if i >= m {
			prePreA[i] = preA[i] + prePreA[i - m]
		} else {
			prePreA[i] = preA[i]
		}
	}

	for i := 1; i < len(prePreA); i++ {
		fmt.Print(prePreA[i], " ")
	}

}