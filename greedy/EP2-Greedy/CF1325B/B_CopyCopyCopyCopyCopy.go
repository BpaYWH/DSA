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

	numCases := nextInt(scn)
	for i := 0; i < numCases; i++ {
		arrLen := nextInt(scn)
		arr := make([]int, arrLen)
		for j := 0; j < arrLen; j++ {
			arr[j] = nextInt(scn)
		}

		// solve here
		solve(arr)
	}

}

func solve(arr []int) {
	set := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		set[arr[i]] = true
	}

	fmt.Println(len(set))
}