package array

func rotate(nums []int, k int)  {

	l := len(nums)

	if l == 1 {
		return
	}

	if k > l {
		k = k%l
	}

	fL := l - k

	s := append([]int{}, nums[fL:]...)

	for i:= 0; i < fL; i++ {
		nums[l- i - 1] = nums[fL - i - 1]
	}

	for j := 0; j < k; j ++ {
		nums[j] = s[j]
	}
}
