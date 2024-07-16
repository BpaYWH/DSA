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

	numCases := nextInt(scn)
	
	for i := 0; i < numCases; i++ {
		numStrings := nextInt(scn)
		strings := make([]string, numStrings)

		for j := 0; j < numStrings; j++ {
			strings[j] = next(scn)
		}

		solve(strings)
	}

}

func solve(strs []string) {
	res := 0
	ones := 0
	zeros := 0
	lens := make([]int, len(strs))

	for i := 0; i < len(strs); i++ {
		lens[i] = len(strs[i])

		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '1' {
				ones++
			} else {
				zeros++
			}	
		}
	}

	for i := 0; i < len(lens); i++ {
		if canBePalindrome(&ones, &zeros, lens[i]) {
			res++
		}
	}

	fmt.Println(res)
}

func canBePalindrome(ones *int, zeros *int, len int) bool {
	newOnes := *ones
	newZeros := *zeros

	if len % 2 == 1 {
		len--
	}
	if *ones % 2 == 1 {
		newOnes--
	}
	if *zeros % 2 == 1 {
		newZeros--
	}

	if newOnes + newZeros >= len {
		if newOnes >= len {
			*ones -= len
			return true
		}

		if newZeros >= len {
			*zeros -= len
			return true
		}

		if len > newOnes{
			*ones -= newOnes
			len -= newOnes
			*zeros -= len
			return true
		}

		if len > newZeros {
			*zeros -= newZeros
			len -= newZeros
			*ones -= len
			return true
		}
	}

	return false
}