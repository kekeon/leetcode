package array

func maxArea(height []int) int {

	i := 0
	l := len(height) - 1
	max := 0

	for l > i {

		if height[i] > height[l] {

			max = Max( height[l] * (l - i), max)
			l --

		} else {
			max = Max( height[i] * (l - i), max)
			i++
		}

	}

	return max

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
