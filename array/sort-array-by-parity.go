package array


func sortArrayByParity(A []int) []int {


	if len(A) == 1 {
		return A
	}

	res1 := []int{}
	res2 := []int{}

	for _, v := range A {
		if isEven(v) {
			res2 = append(res2, v)
		}else {
			res1 = append(res1, v)
		}

	}
	return append(res2, res1...)
}

func isEven(num int) bool{

	if num % 2 == 0 {
		return true
	}

	return false

}