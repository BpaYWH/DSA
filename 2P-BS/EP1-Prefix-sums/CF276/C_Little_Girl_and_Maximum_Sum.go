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

	numEles := nextInt(scn)
	numQueries := nextInt(scn)
	x := make([]int, numEles)
	pre := make([]int, numEles)
	res := 0

	for i := 0; i < numEles; i++ {
		x[i] = nextInt(scn)
	}
	sort.Ints(x)

	for i := 0; i < numQueries; i++ {
		l := nextInt(scn)
		r := nextInt(scn)
		l--
		r--

		pre[l]++
		if r + 1 < numEles {
			pre[r+1]--
		}
	}

	for i := 1; i < len(pre); i++ {
		pre[i] += pre[i - 1]
	}
	sort.Ints(pre)

	for i := 0; i < len(pre); i++ {
		res += pre[i] * x[i]
	}

	fmt.Println(res)
}
