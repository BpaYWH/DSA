package utils

import (
	"math"
)

func MinInt (nums ...int) int {
	minInt := nums[0]

	for _, num := range nums {
		minInt = int(math.Min(float64(minInt), float64(num)))
	}

	return minInt
}

func MaxInt (nums ...int) int {
	maxInt := nums[0]

	for _, num := range nums {
		maxInt = int(math.Max(float64(maxInt), float64(num)))
	}

	return maxInt
}
