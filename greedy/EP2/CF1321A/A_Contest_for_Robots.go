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

	n := nextInt(scn)
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt(scn)
	}
	for i := 0; i < n; i++ {
		y[i] = nextInt(scn)
	}

	// solve here
	solve(x, y)
}

func solve(scoresR []int, scoresB []int) {
	res := 0
	rOnly := 0
	rAndB := 0
	scoreB := 0

	for i := 0; i < len(scoresR); i++ {
		if scoresR[i] > scoresB[i] {
			rOnly++
		}

		if scoresR[i] < scoresB[i] {
			scoreB++
		}

		if scoresR[i] == scoresB[i] {
			rAndB++
			scoreB++
		}
	}

	if (rOnly == 0) {
		fmt.Println(-1)
		return
	}

	diff := scoreB - rAndB
	if (diff % rOnly == 0) {
		res = diff / rOnly + 1
	} else {
		res = int(math.Ceil(float64(diff) / float64(rOnly)))
	}

	fmt.Println(res)
}