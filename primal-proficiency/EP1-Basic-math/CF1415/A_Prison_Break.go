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
		n := nextInt(scn)
		m := nextInt(scn)
		r := nextInt(scn)
		c := nextInt(scn)

		solve(n, m, r, c)
	}
}

func solve(n, m, r, c int) {
	res := 0

	res += int(math.Max(float64(r - 1), float64(n - r)))
	res += int(math.Max(float64(c - 1), float64(m - c)))

	fmt.Println(res)
}