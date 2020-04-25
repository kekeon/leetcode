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
