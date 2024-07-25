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
		x := make([]int, n)
		
		for j := 0; j < n; j++ {
			x[j] = nextInt(scn)
		}

		solve(x)
	}
}

func solve(arr []int) {
	res := 0
	freq := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		diff := arr[i] - i
		if val, ok := freq[diff]; ok {
			res += val
			freq[diff]++
		} else {
			freq[diff] = 1
		}
	}

	fmt.Println(res)
}