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

	numQueries := nextInt(scn)

	for i := 0; i < numQueries; i++ {
		numCandies := nextInt(scn)
		candies := make([]int, numCandies)

		for j := 0; j < numCandies; j++ {
			candies[j] = nextInt(scn)
		}

		solve(candies)
	}
}

func solve(arr []int) {
	res := 0
	freqMap := make(map[int]int)
	sizeArr := []int{}

	for i := 0; i < len(arr); i++ {
		if _, ok := freqMap[arr[i]]; ok {
			freqMap[arr[i]]++
		} else {
			freqMap[arr[i]] = 1
		}
	}

	for _, freq := range freqMap {
		sizeArr = append(sizeArr, freq)
	}

	sort.Ints(sizeArr)
	prev := sizeArr[len(sizeArr)-1]
	res += prev

	for i := len(sizeArr) - 2; i >= 0; i-- {
		if (sizeArr[i] >= prev) {
			res += prev - 1
			prev--
		} else {
			res += sizeArr[i]
			prev = sizeArr[i]
		}
		if prev <= 0 {
			break
		}
	}

	fmt.Println(res)
}