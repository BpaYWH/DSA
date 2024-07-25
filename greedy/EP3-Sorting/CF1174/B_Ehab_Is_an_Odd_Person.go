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
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt(scn)
	}

	// solve here
	solve(x)
}

func solve(arr []int) {
	// check if all numbers are odd
	isAllOdd := true
	isAllEven := true
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 0 {
			isAllOdd = false
			break;
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 1 {
			isAllEven = false
			break;
		}
	}

	if isAllOdd || isAllEven {
		for i := 0; i < len(arr); i++ {
			fmt.Print(arr[i], " ")
		}
		return
	}

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

}