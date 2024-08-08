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

	numCand := nextInt(scn)
	weights := make([]int, numCand)
	for i := 0; i < numCand; i++ {
		weights[i] = nextInt(scn)
	}

	// solve here
	solve(weights)
}

func solve(arr []int) {
	res := 0
	odd := []int{0}
	even := []int{0}

	for i, val := range arr {
		if i % 2 == 0 {
			even = append(even, even[len(even) - 1] + val)
		} else {
			odd = append(odd, odd[len(odd) - 1] + val)
		}
	}

	for i := 0; i < len(arr); i++ {
		evenn := 0
		oddn := 0

		if i % 2 == 0 {
			evenr := even[len(even) - 1] - even[i / 2 + 1]
			oddr := odd[len(odd) - 1] - odd[i / 2]
			evenn = oddr + even[i / 2]
			oddn = evenr + odd[i / 2]
		} else {
			oddr := odd[len(odd) - 1] - odd[(i - 1) / 2 + 1]
			evenr := even[len(even) - 1] - even[(i - 1) / 2 + 1]
			oddn = evenr + odd[(i - 1) / 2]
			evenn = oddr + even[(i - 1) / 2 + 1]
		}

		if evenn == oddn {
			res++
		}
	}


	fmt.Println(res)
}