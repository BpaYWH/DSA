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
	m := nextInt(scn)
	x := make([]string, n)
	y := make([]int, m)
	for i := 0; i < n; i++ {
		x[i] = next(scn)
	}
	for i := 0; i < m; i++ {
		y[i] = nextInt(scn)
	}

	// solve here
	fmt.Println(solve(x, y))
}

func solve(answers []string, points []int) int {
	total := 0

	for i := 0; i < len(points); i++ {
		freqs := make([]int, 5)
		maxCountPos := 0

		for j := 0; j < len(answers); j++ {
			ansPos := rune(answers[j][i]) - rune('A')
			freqs[ansPos]++

			if freqs[maxCountPos] < freqs[ansPos] {
				maxCountPos = int(ansPos)
			}
		}

		total += points[i] * freqs[maxCountPos]
	}

	return total
}
