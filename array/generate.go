package array

func generate(numRows int) [][]int {

	result := [][]int{}
	list := []int{}
	p := 1
	for i := 0; i < numRows; i++ {
		list = []int{}
		for j := 0; j <= i; j++ {
			p = 1
			if j == 0 {
				list = append(list, p)
			} else if j == i {
				list = append(list, p)
			} else {
				p = result[i-1][j-1] + result[i-1][j]
				list = append(list, p)
			}
		}
		result = append(result, list)
	}

	return result
}
