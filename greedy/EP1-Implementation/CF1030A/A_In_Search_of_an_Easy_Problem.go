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
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt(scn)
	}

	// solve here
	res := solve(x)

	fmt.Println(res)
}

func solve(n []int) string {
	for _, cand := range n {
		if cand == 1 {
			return "HARD"
		}
	}

	return "EASY"
}
