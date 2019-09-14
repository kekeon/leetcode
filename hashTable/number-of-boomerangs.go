package hashTable

func numberOfBoomerangs(points [][]int) int {

	res := 0
	for _, v := range points {
		extentMap := map[int]int{}
		for _, v1 := range points {
			if v[0] != v1[0] || v[1] != v1[1] {
				extentMap[pointExtent(v, v1)]++
			}
		}

		for _, v3 := range extentMap {
			res += v3 * (v3 - 1)
		}

	}

	return res
}

func pointExtent(a, b []int) int {

	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}
