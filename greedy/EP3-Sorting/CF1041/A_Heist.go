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

	numRemaining := nextInt(scn)
	keyboards := make([]int, numRemaining)
	for i := 0; i < numRemaining; i++ {
		keyboards[i] = nextInt(scn)
	}

	// solve here
	solve(keyboards)

}

func solve(arr []int) {
	res := 0
	sort.Ints(arr)
	prev := arr[0] - 1

	for i := 0; i < len(arr); i++ {
		if (arr[i] - prev > 1) {
			res += arr[i] - prev - 1
		}
		prev = arr[i]
	}


	fmt.Println(res)
}