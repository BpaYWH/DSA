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

	numRequests := nextInt(scn)
	requests := make([]string, numRequests)
	for i := 0; i < numRequests; i++ {
		requests[i] = next(scn)
	}

	// solve here
	solve(requests)
}

func solve(strs []string) {
	freq := make(map[string]int)

	for i := 0; i < len(strs); i++ {
		if _, ok := freq[strs[i]]; ok {
			freq[strs[i]]++
			fmt.Println(strs[i] + "" + strconv.Itoa(freq[strs[i]]))
		} else {
			freq[strs[i]] = 0
			fmt.Println("OK")
		}
	}
}