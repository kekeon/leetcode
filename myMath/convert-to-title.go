package myMath

func convertToTitle(n int) string {

	m := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	i := 0

	s := ""
	for n > 0 {
		i = ( n-1 ) % 26
		s = m[i] + s
		n = (n -1) / 26
	}

	return s
}
