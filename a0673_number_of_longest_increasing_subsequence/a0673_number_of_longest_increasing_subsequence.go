package a0673_number_of_longest_increasing_subsequence

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	ans := 0
	maxLen := 0
	lens := make([]int, n)
	cnts := make([]int, n)

	for i := 0; i < n; i++ {
		lens[i] = 1
		cnts[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if lens[i] == lens[j]+1 {
					cnts[i] += cnts[j]
				}
				if lens[i] < lens[j]+1 {
					lens[i] = lens[j] + 1
					cnts[i] = cnts[j]
				}
			}
		}
		if maxLen == lens[i] {
			ans += cnts[i]
		}
		if maxLen < lens[i] {
			maxLen = lens[i]
			ans = cnts[i]
		}
	}
	return ans
}
