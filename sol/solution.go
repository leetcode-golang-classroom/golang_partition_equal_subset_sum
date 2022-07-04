package sol

func canPartition(nums []int) bool {
	target_sum := 0
	for _, num := range nums {
		target_sum += num
	}
	if target_sum%2 != 0 {
		return false
	}
	target_sum /= 2
	dp := make(map[int]struct{})
	dp[0] = struct{}{}
	nLen := len(nums)
	for start := nLen - 1; start >= 0; start-- {
		nextDp := make(map[int]struct{})
		for sum := range dp {
			if sum+nums[start] == target_sum {
				return true
			}
			nextDp[target_sum+nums[start]] = struct{}{}
			nextDp[sum] = struct{}{}
		}
		dp = nextDp
	}
	_, exists := dp[target_sum]
	return exists
}
