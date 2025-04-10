package fake

import (
	"math"
	"math/rand/v2"
)

func RandomInt() int {
	return rand.IntN(math.MaxInt)
}

func RandomInts(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = RandomInt()
	}
	return nums
}
