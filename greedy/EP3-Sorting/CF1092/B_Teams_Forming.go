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

	numStudents := nextInt(scn)
	students := make([]int, numStudents)
	for i := 0; i < numStudents; i++ {
		students[i] = nextInt(scn)
	}

	// solve here
	solve(students)
}

func solve(arr []int) {
	res := 0
	sort.Ints(arr)

	for i := 1; i < len(arr); i += 2 {
		res += arr[i] - arr[i-1]
	}

	fmt.Println(res)
}