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
	v := obj.SumRange(0, 5)

	fmt.Println(v)
}

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 4}
	v := lengthOfLIS(nums)
	fmt.Println(v)
}

func TestLongestPalindromeSubseq(t *testing.T) {
	s := "aaa"
	v := longestPalindromeSubseq(s)
	fmt.Println(v)
}

func TestUpdateMatrix(t *testing.T) {
	arr := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	v := updateMatrix(arr)
	fmt.Println(v)
}

func TestMaxProduct(t *testing.T) {
	arr := []int{2, 3, -2, 4}
	v := maxProduct(arr)
	fmt.Println(v)
}


func TestRob(t *testing.T) {
	arr := []int{2,7,9,3,1}
	v := rob2(arr)
	fmt.Println(v)
}

func TestRobTwo(t *testing.T) {
	arr := []int{1, 2, 3, 1}
	v := robTwo2(arr)
	fmt.Println(v)
}
