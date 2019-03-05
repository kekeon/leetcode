package main

import (
	"sort"
	"fmt"
)

type KthLargest struct {
	list []int
	i int
}

func ConstructorFlow(k int, nums []int) KthLargest {
	sort.Ints(nums)
	fmt.Println(nums)
	return KthLargest {
		list: nums,
		i: k,
	}
}

/*func (this *KthLargest) Add(val int) int {

}*/

func main() {
	obj := ConstructorFlow(3, []int{4,5,8,2})

	fmt.Println(obj)
}
