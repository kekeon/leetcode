package array

func checkPossibility(nums []int) bool {
	 n := 0
	 l := len(nums)

	 for i := 0; i< l - 1; i++ {
	 	a := nums[i]
	 	b := nums[i+1]

	 	if a > b {
	 		n++

	 		if n > 1 {
				return false
			}

			if i - 1 > -1 && nums[i-1] > b && i + 2 < l && a > nums[i+2] {
				return false
			}
		}


	 }

	 return true
}