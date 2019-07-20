package bit_manipulation

import (
	"testing"
	"fmt"
)

func TestSingleNumber(t *testing.T) {
	v := singleNumber([]int{4,1,2,1,2,4,3})
	fmt.Println(v)
}
