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

	countA := nextInt(scn)
	countB := nextInt(scn)
	countAB := nextInt(scn)

	// solve here
	solve(countA, countB, countAB)
}

func solve(countA int, countB int, countAB int) {
	res := 2 * (countAB + int(math.Min(float64(countA), float64(countB))))

	if (countA != countB) {
		res += 1
	}

	fmt.Println(res)
}