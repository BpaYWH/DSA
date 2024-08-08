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

	numCards := nextInt(scn)
	cards := make([]int, numCards)
	for i := 0; i < numCards; i++ {
		cards[i] = nextInt(scn)
	}

	// solve here
	solve(cards)
}

func solve(arr []int) {
	res := [2]int{0, 0}
	i := 0
	j := len(arr) - 1
	turnA := 0

	for i <= j {
		if arr[i] < arr[j] {
			res[turnA] += arr[j]
			j--
		} else {
			res[turnA] += arr[i]
			i++
		}
		turnA ^= 1
	}

	fmt.Println(res[0], res[1])
}