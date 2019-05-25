package array

func findMaxAverage(nums []int, k int) float64 {

	l := len(nums)

	if l < k {
		return float64(sumList(nums)) / float64(k)
	}

	sum := sumList(nums[0:k])

	for i:= 1; i <= l - k ; i++ {
		s := sumList(nums[i:i+k])
		if s > sum {
			sum = s
		}
	}

	return  float64(sum) / float64(k)
}

func sumList(nums []int) int{

	s := 0
	for _, v:= range nums {
		s+=v
	}
	return s
}