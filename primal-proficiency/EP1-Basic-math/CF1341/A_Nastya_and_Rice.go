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

	numCasese := nextInt(scn)

	for i := 0; i < numCasese; i++ {
		numGrains := nextInt(scn)
		a := nextInt(scn)
		b := nextInt(scn)
		c := nextInt(scn)
		d := nextInt(scn)

		solve(numGrains, a, b, c, d)
	}
}

func solve(numGrains, a, b, c, d int) {
	minWeight := a - b
	maxWeight := a + b
	minTotalWeights := numGrains * minWeight
	maxTotalWeights := numGrains * maxWeight
	minPossibleWeights := c - d
	maxPossibleWeights := c + d
	if 
		minTotalWeights >= minPossibleWeights && minTotalWeights <= maxPossibleWeights || 
		maxTotalWeights >= minPossibleWeights && maxTotalWeights <= maxPossibleWeights ||
		minTotalWeights <= minPossibleWeights && maxTotalWeights >= maxPossibleWeights ||
		minTotalWeights <= maxPossibleWeights && maxTotalWeights >= minPossibleWeights {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
