package array

func duplicateZeros(arr []int)  {

	cache := []int{}

	for i:=0; i < len(arr); i++ {
		if arr[i] == 0 {
			cache = append(cache, arr[i])
		}

		if len(cache) != 0 {
			cache = append(cache, arr[i])
			arr[i] = cache[0]

			if len(cache) > 1 {
				cache = cache[1:]
			}else {
				cache = []int{}
			}
		}
	}
}
