package array

func twoSum(nums []int, target int) []int {

	m0 := map[int]int{}
	m1 := map[int]int{}

	for i, v := range nums{

		if m0[v] > 0 {
			m1[v] = i + 1
		}else {
			m0[v] = i + 1
		}
	}

	for k, v := range m0{
		if m0[target - k] > 0 {

			// 倍数关系，看是否存在相同元素
			if target - 2 * k == 0 && m1[target - k] > 0{
				return []int{v -1 , m1[target - k] -1}
			}
			// 非倍数关系
			if target - 2 * k != 0 && m0[target - k] > 0{
				return []int{v -1 , m0[target - k] -1}
			}

		}
	}

	return  []int{}
}
