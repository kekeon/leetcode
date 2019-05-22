package array


func matrixReshape(nums [][]int, r int, c int) [][]int {

	eleL := r * c

	list := []int{}

	for i:=0; i < len(nums); i++ {
		list = append(list, nums[i]...)
	}

	ll := len(list)

	if ll != eleL {
		return nums
	}

	res := [][]int{}

	for row := 0 ; row < r; row ++ {
		colList := append([]int{}, list[row*c:(row+1) * c]...)
		res = append(res, colList)
	}

	return res



}
