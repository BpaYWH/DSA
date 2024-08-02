package main

import (
	"bufio"
	"fmt"
	"math"
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
		s := next(scn)

		solve(s)
	}
}

func solve(s string) {
	res := len(s)

	suf0, suf1 := 0, 0
	pre0, pre1 := 0, 0

	for _, c := range s {
		if c == '0' {
			suf0++
		} else {
			suf1++
		}
	}
	res = int(math.Min(float64(suf0), float64(suf1)))

	for _, c := range s {
		if c == '0' {
			pre0++
			suf0--
		} else {
			pre1++
			suf1--
		}

		res = int(math.Min(float64(res), float64(pre0 + suf1)))
		res = int(math.Min(float64(res), float64(pre1 + suf0)))
	}

	fmt.Println(res)
}