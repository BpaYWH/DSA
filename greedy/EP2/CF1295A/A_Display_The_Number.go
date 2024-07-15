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
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	numCases := nextInt(scn)
	for i := 0; i < numCases; i++ {
		fmt.Println(solve(nextInt(scn)))
	}
}

// n: max no. of segments can be turn on
func solve(n int) string {
	numDigits := int(math.Floor(float64(n / 2)))
	res := ""

	for i := 0; i < numDigits; i++ {
		res += "1"

		if i == 0 && n%2 == 1 {
			res = "7"
		}
	}

	return res
}
