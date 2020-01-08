package dp

/**
516. 最长回文子序列
 "bbbab" -> 4
 */
func longestPalindromeSubseq(s string) int {

	l := len(s)
	if l <= 1 {
		return l
	}

	dp := make([][]int, l)

	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		dp[i][i] = 1
	}

	for a := l - 1; a >= 0; a-- {
		for b := a + 1; b < l; b++ {
			if s[a] == s[b] {
				dp[a][b] = dp[a+1][b-1] + 2
			} else {
				dp[a][b] = max(dp[a+1][b], dp[a][b-1])
			}
		}
	}

	return dp[0][l-1]

}
