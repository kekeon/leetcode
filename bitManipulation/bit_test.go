package bitManipulation

import (
	"testing"
	"fmt"
)

func TestSingleNumber(t *testing.T) {
	v := singleNumber([]int{4,1,2,1,2,4,3})
	fmt.Println(v)
}

func TestReverseBits( t *testing.T) {

	v := reverseBits( 43261596)
	fmt.Println(v)
}

func TestToHex( t *testing.T) {

	v := toHex( -2)
	fmt.Println(v)
}