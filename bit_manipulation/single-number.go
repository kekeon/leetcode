package bit_manipulation

func singleNumber(nums []int) int {
	res := 0
	for _,v := range nums {
		res = res ^ v
	}
	return res
}