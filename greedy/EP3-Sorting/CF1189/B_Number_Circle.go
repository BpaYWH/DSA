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

	n := nextInt(scn)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt(scn)
	}

	// solve here
	solve(x)
}

func solve(arr []int) {
	sort.Ints(arr)

	target := arr[len(arr) - 1] - arr[len(arr) - 2]

	for i := 0; i < len(arr) - 2; i++ {
		if arr[i] > target {
			fmt.Println("YES")

			fmt.Print(arr[i], " ")
			for j := 0; j < len(arr); j++ {
				if j != i {
					fmt.Print(arr[j], " ")
				}
			}

			return
		}
	}

	fmt.Println("NO")
}