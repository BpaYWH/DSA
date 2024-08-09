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

	nextInt(scn)
	quota := nextInt(scn)
	str := next(scn)

	// solve here
	solve(str, quota)
}

func solve(str string, quota int) {
	res := 0

	res = int(math.Max(float64(walkString(str, quota, 'a')), float64(walkString(str, quota, 'b'))))

	fmt.Println(res)
}

func walkString(str string, quota int, t rune) int {
	res := 0
	q := &Queue[int]{}
	i := 0
	j := 0
	cnt := 0

	for id, val := range str {
		if val == t {
			cnt++
		} else {
			if q.Size() < quota {
				cnt++
				q.Enqueue(id)
			} else {
				next, hasNext := q.Dequeue()
				if !hasNext {
					i = j + 1
					cnt = 0
				} else {
					cnt -= next - i
					i = next + 1
					q.Enqueue(j)
				}
			}
		}

		j++
		res = int(math.Max(float64(res), float64(cnt)))
	}

	return res
}



type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T

	if len(q.items) == 0 {
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) Front() (T, bool) {
	var zero T

	if len(q.items) == 0 {
		return zero, false
	}

	return q.items[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}