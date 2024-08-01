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

	n := nextInt(scn)
	m := nextInt(scn)

	solve(n, m)
}

// (a + b) % c = ((a % c) + (b % c)) % c
func solve(n, m int) {
	res := 0

	for i := 1; i <= m; i++ {
		res += int(math.Floor(float64((n + (i %  5)) / 5)))
	}

	fmt.Println(res)
}