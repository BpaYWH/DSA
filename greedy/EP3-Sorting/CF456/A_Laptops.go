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

	numLaptops := nextInt(scn)
	laptops := make([][2]int, numLaptops)
	for i := 0; i < numLaptops; i++ {
		laptops[i][0] = nextInt(scn)
		laptops[i][1] = nextInt(scn)
	}

	// solve here
	solve(laptops)
}

func solve(arr [][2]int) {
	sort.Slice(arr, func(i, j int) bool {
        if arr[i][0] != arr[j][0] {
            return arr[i][0] < arr[j][0]
        }
        return arr[i][1] < arr[j][1]
    })
	
	maxQuality := arr[0][1]
	for i := 1; i < len(arr); i++ {
		if arr[i][1] < maxQuality {
			fmt.Println("Happy Alex")
			return
		}
		maxQuality = max(maxQuality, arr[i][1])
	}

	fmt.Println("Poor Alex")
}