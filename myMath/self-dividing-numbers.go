package myMath
func selfDividingNumbers(left int, right int) []int {

	arr := []int{}
	var a int

	for i := left; i <= right; i++ {
		num := i

		for num != 0 {
			a = num % 10
			if a == 0 || i % a != 0{
				break
			}else {
				num /= 10
			}
		}

		if num == 0 {
			arr = append(arr, i)
		}
	}

	return arr

}
