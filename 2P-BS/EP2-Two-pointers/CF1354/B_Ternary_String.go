package main

import (
	"bufio"
	"fmt"
	"math"
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

	numCases := nextInt(scn)
	for i := 0; i < numCases; i++ {
		s := next(scn)
		solve(s)
	}
}

func solve(str string) {
	res := len(str) + 1
	i := 0
	j := 0
	freq := make(map[rune] int)
	
	for _, c := range str {
		if _, ok := freq[c]; ok {
			freq[c]++
		} else {
			freq[c] = 1
		}
		
		for i < j && len(freq) == 3 {
			// fmt.Println(i, j)
			res = MinInt(res, j - i + 1)
			val := rune(str[i])
			if freq[val] > 1 {
				freq[val]--
				i++
			} else {
				break
			}
		}

		j++
	}
	
	fmt.Println(res % (len(str) + 1))
}

func MinInt (nums ...int) int {
	minInt := nums[0]

	for _, num := range nums {
		minInt = int(math.Min(float64(minInt), float64(num)))
	}

	return minInt
}

func MaxInt (nums ...int) int {
	maxInt := nums[0]

	for _, num := range nums {
		maxInt = int(math.Max(float64(maxInt), float64(num)))
	}

	return maxInt
}