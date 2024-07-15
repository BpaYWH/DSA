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

	numCases := nextInt(scn)

	for i := 0; i < numCases; i++ {
		size := nextInt(scn)
		arr := make([]int, size)

		for j := 0; j < size; j++ {
			arr[j] = nextInt(scn)
		}

		solve(arr)
	}

}

func solve(arr []int) {
	size := len(arr)
	sort.Ints(arr)

	res := arr[size-1] * arr[size-2] * arr[size-3] * arr[size-4] * arr[size-5]
	res = max(res, arr[0] * arr[size-1] * arr[size-2] * arr[size-3] * arr[size-4])
	res = max(res, arr[0] * arr[1] * arr[size-1] * arr[size-2] * arr[size-3])
	res = max(res, arr[0] * arr[1] * arr[2] * arr[size-1] * arr[size-2])
	res = max(res, arr[0] * arr[1] * arr[2] * arr[3] * arr[size-1])
	res = max(res, arr[0] * arr[1] * arr[2] * arr[3] * arr[4])

	fmt.Println(res)
}