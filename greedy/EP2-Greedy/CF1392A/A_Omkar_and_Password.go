package main

import (
	"bufio"
	"container/list"
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
		lenPw := nextInt(scn)
		pw := make([]int, lenPw)
		for j := range pw {
			pw[j] = nextInt(scn)
		}
		// solve here
		fmt.Println(solve(pw))
	}
}

func solve(pw []int) int {
	if len(pw) <= 1 {
		return len(pw)
	}
	if len(pw) == 2 {
		if pw[0] != pw[1] {
			return 1
		} else {
			return 2
		}
	}

	l := list.New()
	for i := range pw {
		l.PushBack(pw[i])
	}

	e := l.Front()
	for e.Next() != nil && e.Next().Next() != nil {
		e1 := e.Next()
		e2 := e.Next().Next()

		if e.Value.(int) != e1.Value.(int) && e.Value.(int)+e1.Value.(int) != e2.Value.(int) {
			e.Value = e.Value.(int) + e1.Value.(int)
			l.Remove(e1)
			continue
		}

		if e1.Value.(int) != e2.Value.(int) {
			e1.Value = e1.Value.(int) + e2.Value.(int)
			l.Remove(e2)
			continue
		}

		e = e.Next()
	}

	e = l.Back()
	for e.Prev() != nil && e.Prev().Prev() != nil {
		e1 := e.Prev()
		e2 := e.Prev().Prev()

		if e.Value.(int) != e1.Value.(int) && e.Value.(int)+e1.Value.(int) != e2.Value.(int) {
			e.Value = e.Value.(int) + e1.Value.(int)
			l.Remove(e1)
			continue
		}

		if e1.Value.(int) != e2.Value.(int) {
			e1.Value = e1.Value.(int) + e2.Value.(int)
			l.Remove(e2)
			continue
		}

		e = e.Prev()
	}

	if l.Len() == 2 {
		if l.Front().Value.(int) == l.Back().Value.(int) {
			return 2
		} else {
			return 1
		}
	}

	return l.Len()
}

/*
Time: O(n)
Space: O(n)
*/
