package array

import (
	"testing"
	"fmt"
)

// go test -v .\array -test.run TestMaxSubArray
func TestMaxSubArray(t *testing.T){
	nums := []int{2,1}
	v := maxSubArray(nums)
	fmt.Println(v)
}


// go test -v .\array -test.run TestGenerate
func TestGenerate(t *testing.T){
	v := generate(5)
	fmt.Println(v)
}