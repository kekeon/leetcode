package hash_table

func isIsomorphic(s string, t string) bool {
	sm := map[string]int{}
	tm := map[string]int{}

	for i:=0; i < len(s); i++ {
		sm[string(s[i])] += i + 1
		tm[string(t[i])] += i + 1


		if sm[string(s[i])] != tm[string(t[i])] {
			return false
		}
	}

	return true
}
