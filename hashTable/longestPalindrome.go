package hashTable

func longestPalindrome(s string) int {

	m := map[string]int{}
	oddNum := 0
	count := 0
	for _, v := range s {
		m[string(v)] ++
	}

	for _, v := range m {
		if v%2 == 0 {
			count += v
		} else if oddNum > 0 {
			oddNum += v - 1
		} else {
			oddNum += v
		}
	}

	return count + oddNum
}
