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

	n := nextInt(scn)

	// solve here
	solve(n)
}

func solve(n int) {
	for i := 1; i <= n; i++ {
		step := i
		for j := 0; j < n; j++ {
			if  j % 2 == 1 {
				fmt.Print(step, " ")
				step += 2 * (i - 1)
			} else {
				fmt.Print(step , " ")
				step += 2 * (n - i)
			}
			step++
		}
		fmt.Println()
	}	
}