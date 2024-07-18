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

	s := next(scn)

	// solve here
	solve(s)
}

func solve(str string) {
	arr := make([]int, 0)
	for i := 0; i < len(str); i++ {
		if str[i] != '+' {
			num, _ := strconv.Atoi(string(str[i]))
			arr = append(arr, num)
		}
	}

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])

		if i != len(arr) - 1 {
			fmt.Print("+")
		}
	}
}