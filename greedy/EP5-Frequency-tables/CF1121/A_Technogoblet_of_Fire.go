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

	numStudents := nextInt(scn)
	numSchools := nextInt(scn)
	numChosen := nextInt(scn)
	powStudents := make([]int, numStudents)
	fromSchools := make([]int, numStudents)
	idChosens := make([]int, numChosen)

	for i := 0; i < numStudents; i++ {
		powStudents[i] = nextInt(scn)
	}
	for i := 0; i < numStudents; i++ {
		fromSchools[i] = nextInt(scn)
	}
	for i := 0; i < numChosen; i++ {
		idChosens[i] = nextInt(scn)
	}

	// solve here
	solve(numSchools, powStudents, fromSchools, idChosens)
}

func solve(numSchools int, powers []int, fromSchools []int, idChosens []int) {
	res := 0
	powMap := make(map[int](map[string]int)) // O(m)
	idSet := make(map[int]bool) // O(k)
	maxSet := make(map[int]bool) // O(m)

	// O(n)
	for i := 0; i < len(fromSchools); i++ {
		powMap[fromSchools[i]] = map[string]int{"power": 0, "id": 0}
	}

	// O(n)
	for i := 0; i < len(powers); i++ {
		if powers[i] > powMap[fromSchools[i]]["power"] {
			powMap[fromSchools[i]]["power"] = powers[i]
			powMap[fromSchools[i]]["id"] = i + 1
		}
	}

	// O(k)
	for i := 0; i < len(idChosens); i++ {
		idSet[idChosens[i]] = true
	}

	// O(m)
	for _, info := range powMap {
		maxSet[info["id"]] = true
	}

	// O(k)
	for id, _ := range idSet {
		if !maxSet[id] {
			res++
		}
	}

	fmt.Println(res)
}