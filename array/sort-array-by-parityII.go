package array

func sortArrayByParityII(A []int) []int {

	l := len(A)
	oddList := []int{}
	evenList := []int{}

	for i:=0; i < l; i++ {
		if A[i]%2 == 0 {
			evenList = append(evenList, A[i])
		}else {
			oddList = append(oddList, A[i])
		}
	}

	for j:=0; j < l; j++ {
		if j%2 == 0 {
			A[j] = evenList[j/2]
		}else {
			A[j] = oddList[(j-1) / 2]
		}
	}

	return A

}
