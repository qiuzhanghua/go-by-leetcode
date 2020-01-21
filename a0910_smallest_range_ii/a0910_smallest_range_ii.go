package a0910_smallest_range_ii

import "sort"

func smallestRangeII(A []int, K int) int {
	N := len(A)
	sort.Ints(A)
	ans := A[N-1] - A[0]
	for i := 0; i < N-1; i++ {
		a, b := A[i], A[i+1]
		high := A[N-1] - K
		if high < a+K {
			high = a + K
		}
		low := A[0] + K
		if low > b-K {
			low = b - K
		}
		if ans > high-low {
			ans = high - low
		}
	}
	return ans
}
