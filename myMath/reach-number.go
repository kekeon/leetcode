package myMath

func reachNumber(target int) int {

	t := target

	if t < 0 {
		t = -t
	}

	sum := 0
	f := 0
	for sum < t {
		f++
		sum += f
	}
	if (sum-t)%2 == 0 {
		return f
	}
	return f + 1 + f%2
}
