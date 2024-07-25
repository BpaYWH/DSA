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

	numFlats := nextInt(scn)
	lights := make([]int, numFlats)
	for i := 0; i < numFlats; i++ {
		lights[i] = nextInt(scn)
	}

	// solve here
	solve(lights)
}

func solve(lights []int) {
	res := 0
	prevDis := -1

	for i := 1; i < len(lights) - 1; i++ {
		if (lights[i] == 0 && lights[i-1] == 1 && lights[i+1] == 1) {
			if (prevDis == -1 || i - prevDis > 2) {
				prevDis = i
				res++
			} else {
				prevDis = -1
			}
		}
	}

	fmt.Println(res)
}