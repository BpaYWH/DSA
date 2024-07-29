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

	numQueries := nextInt(scn)
	for i := 0; i < numQueries; i++ {
		lenList := nextInt(scn)
		pos := nextInt(scn)

		solve(lenList, pos)
	}
}

func solve(n, x int) {
	fmt.Println(2 * x)
}