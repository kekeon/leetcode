package main

import "fmt"

type RecentCounter struct {
	ints []int
}


func Constructor() RecentCounter {
	return RecentCounter {
		ints: []int{},
	}
}


func (this *RecentCounter) Ping(t int) int {
	this.ints = append(this.ints, t)
	fmt.Println()
	for this.ints[0] < t-3000{
		this.ints = this.ints[1:]
	}
	
	return  len(this.ints)
}

func main() {

		t := 1
	 obj := Constructor()
	 v := obj.Ping(t)

	fmt.Println(v)
}
