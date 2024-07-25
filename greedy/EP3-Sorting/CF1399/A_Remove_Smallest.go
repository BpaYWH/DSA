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
		lenArr := nextInt(scn)
		arr := make([]int, lenArr)

		for j := 0; j < lenArr; j++ {
			arr[j] = nextInt(scn)	
		}

		solve(arr)
	}
}

func solve(arr []int) {
	sort.Ints(arr)

	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] - prev > 1 {
			fmt.Println("NO")
			return
		}
		prev = arr[i]
	}


	fmt.Println("YES")
}