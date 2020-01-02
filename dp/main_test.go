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

	list1 := [][]int{
		{7, 6, 3},
		{8, 7, 2},
		{7, 3, 1},
	}

	v := minPathSum(list)
	fmt.Println(v)
	fmt.Println(list1)
}
