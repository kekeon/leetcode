package array

func numPairsDivisibleBy60(time []int) int {

	arr := make([]int, 60)
	count := 0
	for _, v := range time{
		if v % 60 == 0 {
			count += arr[0]
		}else {
			count += arr[60 - v % 60]
		}

		arr[v%60]++
	}

	return count
}
