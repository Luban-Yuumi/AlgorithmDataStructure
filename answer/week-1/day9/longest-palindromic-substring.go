package day9

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	maxLen := 0
	begin := 0
	for l := 2; l <= n; l++ {
	loop1:
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j > n-1 {
				goto loop1
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if l == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && l > maxLen {
				maxLen = l
				begin = i
			}
		}
	}
	return s[begin:begin+maxLen]
}
