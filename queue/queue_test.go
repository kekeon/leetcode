package queue

import (
	"testing"
	"fmt"
)

// go test -v queue_test.go -test.run TestRecentCounter
func TestRecentCounter(t *testing.T) {
	p := 1
	obj := Constructor()
	v := obj.Ping(p)
	fmt.Println(v)
}
