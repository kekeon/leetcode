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
