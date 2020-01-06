package dp

import (
	"testing"
	"fmt"
)

func TestClimbStairs(t *testing.T) {
	v := climbStairs(3)
	fmt.Println(v)
}

func TestUniquePaths(t *testing.T) {
	v := uniquePaths(7, 3)
	fmt.Println(v)
}

func TestMinPathSum(t *testing.T) {

	list := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}


	v := minPathSum(list)
	fmt.Println(v)
}

func TestSumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	v := obj.SumRange(0,5)

	fmt.Println(v)
}

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10,9,2,5,3,7,101,18}
	v := lengthOfLIS(nums)
	fmt.Println(v)
}
