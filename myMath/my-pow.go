package myMath

func myPowFloat(x float64, n int) float64 {

	if n >= 0 {
		return pow(x, n)
	} else {
		n = -n
		return 1 / pow(x, n)
	}

}

func pow(x float64, n int) float64 {

	if n == 0 {
		return  1
	}

	if n % 2 == 0 {
		r := pow(x, n / 2)
		return r * r
	} else {
		r := pow(x, (n - 1) / 2)
		return r * r * x
	}
}
