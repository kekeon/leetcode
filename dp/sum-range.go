package dp

type NumArray struct {
	list []int
}

// 303. 区域和检索 - 数组不可变
func Constructor(nums []int) NumArray {
	return NumArray {
		list: nums,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for i <= j {
		sum += this.list[j]
		j--
	}
	return sum
}
