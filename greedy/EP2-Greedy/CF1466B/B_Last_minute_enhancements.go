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
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)

	numCases := nextInt(scn)
	for i := 0; i < numCases; i++ {
		lenSong := nextInt(scn)
		song := make([]int, lenSong)
		for j := 0; j < lenSong; j++ {
			song[j] = nextInt(scn)
		}

		// solve here
		fmt.Println(solve(song))
	}
}

// Time: O(n)
// Space: O(n)
func solve(song []int) int {
	// two hash sets, one for unchanged and one for changed
	unchanged := make(map[int]bool)
	changed := make(map[int]bool)

	// iterate through song
	for i := 0; i < len(song); i++ {
		// add to unchanged set if the number is not in the unchanged set
		// else increment by 1 and add to changed set
		if !unchanged[song[i]] && !changed[song[i]] {
			unchanged[song[i]] = true
		} else {
			changed[song[i]+1] = true
		}
	}

	// fmt.Println(unchanged)
	// fmt.Println(changed)
	// return the sum of the length of the changed set and the length of the unchanged set
	return len(changed) + len(unchanged)
}
