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

	numCases := nextInt(scn)
	for i := 0; i < numCases; i++ {
		numEles := nextInt(scn)
		eles := make([]int, numEles)
		for j := 0; j < numEles; j++ {
			eles[j] = nextInt(scn)
		}

		solve(eles)
	}
}

func solve(eles []int) {
	res := 0

	sign := 1
	if eles[0] < 0 {
		sign = -1
	}
	opt := eles[0]

	for _, val := range eles {
		currSign := 1 
		if val < 0 {
			currSign = -1
		}

		if currSign != sign {
			res += opt
			sign = currSign
			opt = val
		} else {
			opt = int(math.Max(float64(opt), float64(val)))
		}
	}

	res += opt

	fmt.Println(res)
}