package array

import (
	"testing"
	"fmt"
)

// go test -v .\array -test.run TestMaxSubArray
func TestMaxSubArray(t *testing.T) {
	nums := []int{2, 1}
	v := maxSubArray(nums)
	fmt.Println(v)
}

// go test -v .\array -test.run TestGenerate
func TestGenerate(t *testing.T) {
	v := generate(5)
	fmt.Println(v)
}

func TestGetRow(t *testing.T) {
	v := getRow(3)
	fmt.Println(v)
}

func TestMaxProfit(t *testing.T) {
	v := maxProfit([]int{7, 1, 5, 3, 6, 4})

	fmt.Println(v)
}

// [1,2,3,4,4,9,56,90]

func TestTwoSum(t *testing.T) {
	v := twoSum([]int{2, 7, 11, 15}, 18)

	fmt.Println(v)
}

// [1,2,3,4,4,9,56,90]

func TestRotate(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(v, 3)
	fmt.Println(v)
}

func TestContainsNearbyDuplicate(t *testing.T) {
	v := containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
	fmt.Println(v)
}

func TestFindDisappearedNumbers(t *testing.T) {
	v := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(v)
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	v := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})

	fmt.Println(v)
}

func TestFib(t *testing.T) {
	v := fib(4)
	fmt.Println(v)
}

func TestFindPairs(t *testing.T) {
	// v := findPairs([]int{3, 1, 4, 1, 5}, 2)
	v := findPairs([]int{0, 0, 1, 0, 0}, 1)
	fmt.Println(v)
}


func TestArrayPairSum(t *testing.T) {
	// v := findPairs([]int{1,4,3,2})
	v := arrayPairSum([]int{-1,4,3,-2})
	fmt.Println(v)
}

func TestFindUnsortedSubarray(t *testing.T){
	v := findUnsortedSubarray([]int{1,3,2,2,2})
	fmt.Println(v)
}

func TestMaximumProduct(t *testing.T){
	v := maximumProduct([]int{722,634,-504,-379,163,-613,-842,-578,750,951,-158,30,-238,-392,-487,-797,-157,-374,999,-5,-521,-879,-858,382,626,803,-347,903,-205,57,-342,186,-736,17,83,726,-960,343,-984,937,-758,-122,577,-595,-544,-559,903,-183,192,825,368,-674,57,-959,884,29,-681,-339,582,969,-95,-455,-275,205,-548,79,258,35,233,203,20,-936,878,-868,-458,-882,867,-664,-892,-687,322,844,-745,447,-909,-586,69,-88,88,445,-553,-666,130,-640,-918,-7,-420,-368,250,-786})
	fmt.Println(v)
	// 943695360
}