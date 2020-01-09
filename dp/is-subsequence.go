package dp

func isSubsequence(s string, t string) bool {
	subIndex := 0
	subLen := len(s)

	if subLen == 0 {
		return true
	}

	for _, v := range t {
		if string(v) == string(s[subIndex]) {
			subIndex++
		}

		if subIndex > subLen {
			return true
		}
	}

	return false
}

