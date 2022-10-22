package maximum_profit_in_job

import "sort"

func jobScheduling(startTime, endTime, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][1] < jobs[j][1] })

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		k := sort.Search(i, func(j int) bool { return jobs[j][1] > jobs[i-1][0] })
		dp[i] = max(dp[i-1], dp[k]+jobs[i-1][2])
	}
	return dp[n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
