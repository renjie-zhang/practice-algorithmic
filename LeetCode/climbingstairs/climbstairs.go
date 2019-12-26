/*
 * @Author: your name
 * @Date: 2019-12-26 15:06:56
 * @LastEditTime: 2019-12-26 17:46:30
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \GoPath\src\Practice-Algorithmic\LeetCode\climbingstairs\climbstairs.go
 */
package climbingstairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
