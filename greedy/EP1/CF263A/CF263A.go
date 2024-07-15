package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	var mat []string
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 5; i++ {
		if scanner.Scan() {
			row := scanner.Text()
			mat = append(mat, strings.ReplaceAll(row, " ", ""))
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == '1' {
				fmt.Println(int(math.Abs(float64(i-2)) + math.Abs(float64(j-2))))
				break
			}
		}
	}
}
