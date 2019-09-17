package myMath

func convertToTitle(n int) string {

	m := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if n <= 26 {
		return m[n-1]
	}

	b := (n - n % 26) / 26

	r := (n - b - 1) % 25

	if b == 0 {

		return m[b] + m[((n - 1) % 26) - 1]

	} else {
		return m[b - 1] + m[r]
	}

}
