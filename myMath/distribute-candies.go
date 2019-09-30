package myMath

// 1103
func distributeCandies(candies int, num_people int) []int {

	last := 0 // 最一次分配的数量
	arr := make([]int, num_people)

	for candies > 0 {
		for i := 0; i < num_people; i++ {
			if candies > last+1 {
				last ++
				arr[i] = arr[i] + last
				candies = candies - last
			} else {
				arr[i] = arr[i] + candies
				candies = 0
				break
			}

		}
	}

	return arr

}
