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

	numCases := nextInt(scn)

	for i := 0; i < numCases; i++ {
		len := nextInt(scn)
		seq := make([]int, len)
		
		for j := 0; j < len; j++ {
			seq[j] = nextInt(scn)
		}

		solve(seq)
	}
}

func solve(seq []int) {
	i := 0;
	j := len(seq) - 1;
	
	for  i < j {
		fmt.Print(seq[i], " ")
		fmt.Print(seq[j], " ")

		i++
		j--
	}

	if i == j {
		fmt.Print(seq[i])
	}

	fmt.Println()
}