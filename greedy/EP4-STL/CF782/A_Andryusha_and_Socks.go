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

	numSocks := nextInt(scn)
	socks := make([]int, numSocks * 2)
	for i := 0; i < len(socks); i++ {
		socks[i] = nextInt(scn)
	}

	// solve here
	solve(socks)
}

func solve(arr []int) {
	res := 0
	freq := make(map[int]int)
	num := 0

	for i := 0; i < len(arr); i++ {
		if _, ok := freq[arr[i]]; ok {
			num--
		} else {
			num++
			res = max(res, num)
			freq[arr[i]] = 1
		}
	}

	fmt.Println(res)
}