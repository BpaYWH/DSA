package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	_ = nextInt(scn)
	rooms := next(scn)

	// solve here
	solve(rooms)
}

func solve(str string) {
	res := 0
	freq := make(map[string]int)

	for i := 0; i < len(str); i++ {
		if i % 2 == 1 {
			key := strings.ToLower(string(str[i]))

			if val, ok := freq[key]; ok && val > 0 {
				freq[key]--
			} else {
				res++
			}
		} else {
			if _, ok := freq[string(str[i])]; ok {
				freq[string(str[i])]++
			} else {
				freq[string(str[i])] = 1
			}
		}
	}

	fmt.Println(res)
}