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
		numPart := nextInt(scn)
		numProblems := make([]int, numPart)

		for j := 0; j < numPart; j++ {
			numProblems[j] = nextInt(scn)
		}

		solve(numProblems)
	}
}

func solve(arr []int) {
	pre := []int{}
	prev := arr[0]
	cnt := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != prev {
			pre = append(pre, cnt)
			prev = arr[i]
			cnt = 1
		} else {
			cnt++
		}
	}
	pre = append(pre, cnt)

	prepre := make([]int, len(pre) + 1)
	prepre[0] = 0

	for i, val := range pre {
		prepre[i + 1] = val + prepre[i]
	}

	g := prepre[1]
	s := 0
	b := 0

	for i := 2; i < len(prepre) && prepre[i] < len(arr) / 2; i++ {
		if prepre[i] > g * 2 {
			s = prepre[i] - g

			j := i + 1
			for prepre[j] <= len(arr) / 2 {
				j++
			}
			j--

			if j > i && prepre[j] - prepre[i] > g{
				b = prepre[j] - prepre[i]
				break
			}
		}
	}

	if b == 0 {
		fmt.Println("0 0 0")
	} else {
		fmt.Println(g, s, b)
	}
}