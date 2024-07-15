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

	n := nextInt(scn)
	res := "NO"
	// solve here
	res = solve(n)

	fmt.Println(res)
}

func solve(n int) string {
	count := 0

	for n > 0 {
		if n%10 == 4 || n%10 == 7 {
			count++
		}
		n /= 10
	}

	if count == 4 || count == 7 {
		return "YES"
	}

	return "NO"
}
