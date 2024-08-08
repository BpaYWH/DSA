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

	numBars := nextInt(scn)
	bars := make([]int, numBars)
	for i := 0; i < numBars; i++ {
		bars[i] = nextInt(scn)
	}

	// solve here
	solve(bars)
}

func solve(arr []int) {
	res := [2]int{0, 0}

	i := 0
	j := len(arr) - 1
	left := false
	right := false

	for i < j {
		if  arr[i] < arr[j] {
			arr[j] -= arr[i]
			res[0]++
			left = true
			right = false
			i++
			continue
		}
		if arr[i] > arr[j] {
			arr[i] -= arr[j]
			res[1]++
			left = false
			right = true
			j--
			continue
		}
		if arr[i] == arr[j] {
			arr[i] = 0
			arr[j] = 0
			res[0]++
			res[1]++
			left = true
			right = true
			i++
			j--
		}
	}

	if i == j {
		if left && right {
			res[0]++
		} else {
			if left {
				res[1]++
			} else {
				res[0]++
			}
		}
	}

	fmt.Println(res[0], res[1])
}