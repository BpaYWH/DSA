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
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	n1 := nextInt(scn)
	n2 := nextInt(scn)

	// solve here
	fmt.Println(solve(n1, n2))
}

func solve(n1 int, n2 int) int {
	for i := 0; i < n2; i++ {
		if n1%10 != 0 {
			n1 -= 1
		} else {
			n1 /= 10
		}
	}

	return n1
}
