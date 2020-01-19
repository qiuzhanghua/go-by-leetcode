package a1000_minimum_cost_to_merge_stones

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func mergeStones(stones []int, K int) int {
	if K == 0 {
		return 0
	}
	n := len(stones)
	if (n-1)%(K-1) != 0 {
		return -1
	}

	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + stones[i-1]
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for r := 0; r < n; r++ {
		for l := r - K + 1; l >= 0; l-- {
			max := MaxInt
			for p := r; p > l; p -= K - 1 {
				if dp[l][p-1]+dp[p][r] < max {
					max = dp[l][p-1] + dp[p][r]
				}
				dp[l][r] = max
			}
			if (r-l)%(K-1) == 0 {
				dp[l][r] += prefixSum[r+1] - prefixSum[l]
			}
		}
	}
	return dp[0][n-1]
}
