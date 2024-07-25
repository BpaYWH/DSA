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

	numElements := nextInt(scn)
	elements := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		elements[i] = nextInt(scn)
	}

	// solve here
	solve(elements)
}

func solve(arr []int) {
	freqMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if _, ok := freqMap[arr[i]]; ok {
			if (freqMap[arr[i]] >= 2) {
				fmt.Print("NO")
				return
			}
			freqMap[arr[i]]++
		} else {
			freqMap[arr[i]] = 1
		}
	}

	asc := []int{}
	desc := []int{}

	for ele, freq := range freqMap {
		if freq > 1 {
			asc = append(asc, ele)
			desc = append(desc, ele)
		} else {
			desc = append(desc, ele)
		}
	}

	sort.Ints(asc)
	sort.Ints(desc)

	fmt.Println("YES")
	fmt.Println(len(asc))
	for i := 0; i < len(asc); i++ {
		fmt.Print(asc[i], " ")
	}
	fmt.Println()

	fmt.Println(len(desc))
	for i := len(desc) - 1; i >= 0; i-- {
		fmt.Print(desc[i], " ")
	}	
}