package myMath

import (
	"fmt"
	"testing"
)

func TestRecentCounter(t *testing.T) {
	v := mySqrt(36)
	fmt.Println(v)
}

func TestTitleToNumber(t *testing.T) {
	v := titleToNumber2("ZY")
	fmt.Println(v)
}

func TestTrailingZeroes(t *testing.T) {
	v := trailingZeroes(30)
	fmt.Println(v)
}

func TestBinaryGap(t *testing.T) {
	v := binaryGap(5)
	fmt.Println(v)
}

func TestMyPow(t *testing.T) {
	v := myPowFloat(2, 1)
	fmt.Println(v)
}

func TestSubsets(t *testing.T) {
	v := subsets([]int{1, 2, 2})
	fmt.Println(v)

	/*

	func backtrackResultUse(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for a := 0; a <= len(nums); a++ {
		backtrack(a, 0, []int{}, nums, &result)
	}

	return result
}

func backtrackUse(size int, start int, path []int, nums []int, result *[][]int) {
	if size == len(path) {
		p := append([]int{}, path...)
		*result = append(*result, p)
		return
	}

	for a := start; a < len(nums); a++ {

		if a > start && nums[a] == nums[a - 1] {
			continue
		}

		path = append(path, nums[a])
		backtrack(size, a+1, path, nums, result)
		path = path[0 : len(path)-1]
	}
}


	*/
}

func TestCombine(t *testing.T) {
	v := combine(4, 2)
	fmt.Println(v)
}

func TestGrayCode(t *testing.T) {
	v := grayCode(2)
	fmt.Println(v)
}

func TestRestoreIpAddresses(t *testing.T) {
	v := restoreIpAddresses("101023")
	fmt.Println(v, len(v))
}

func TestExist(t *testing.T) {
	b := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	w := "ABCB"
	v := exist(b, w)
	fmt.Println(v, v)
}

func TestSolveNQueens(t *testing.T) {
	v := solveNQueens(4)
	fmt.Println(v)
}
