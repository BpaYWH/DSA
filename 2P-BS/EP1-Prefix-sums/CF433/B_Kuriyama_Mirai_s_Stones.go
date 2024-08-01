package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	y := make([]int, n)
	preX := make([]int, n)
	preY := make([]int, n)
	for i := 0; i < n; i++ {
		y[i] = nextInt(scn)
		if i > 0 {
			preX[i] = preX[i - 1] + y[i]
		} else {
			preX[i] = y[i]
		}
	}

	sort.Ints(y)
	preY[0] = y[0]

	for i := 1; i < n; i++ {
		preY[i] = preY[i - 1] + y[i]
	}


	m := nextInt(scn)
	for i := 0; i < m; i++ {
		t := nextInt(scn)
		l := nextInt(scn)
		r := nextInt(scn)
		if (t != 1) {
			solve(preY, l, r)
		} else {
			solve(preX, l, r)
		}
	}


}

func solve(arr []int, l, r int) {
	res := 0
	res = arr[r - 1]

	if (l >= 2) {
		res -= arr[l - 2]
	} 

	fmt.Println(res)
}