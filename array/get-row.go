package array

func getRow(rowIndex int) []int {

	rest := make([]int, rowIndex + 1)
	rest[0] = 1
	for i:=1;i<=rowIndex;i++ {
		for j:=i; j >0 ; j-- {
			rest[j] += rest[j-1]
		}
	}

	return rest
}
