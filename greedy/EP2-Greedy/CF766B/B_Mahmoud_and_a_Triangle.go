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
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	n := nextInt(scn)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt(scn)
	}

	// solve here
	sorting(x)
}

// sorting with sliding window, O(n log n)
func sorting(x []int) {
	sort.Ints(x)

	// sliding window of size 3
	for i := 0; i < len(x)-2; i++ {
		if isValidTriangle(x[i], x[i+1], x[i+2]) {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}

// brute force, O(n^3)
func bruteForce(x []int) {
	for i := 0; i < len(x)-2; i++ {
		for j := i + 1; j < len(x)-1; j++ {
			for k := j + 1; k < len(x); k++ {
				if isValidTriangle(x[i], x[j], x[k]) {
					fmt.Println("YES")
					return
				}
			}
		}
	}

	fmt.Println("NO")
}

func isValidTriangle(x1 int, x2 int, x3 int) bool {
	return x1 + x2 > x3 && x1 + x3 > x2 && x2 + x3 > x1
}