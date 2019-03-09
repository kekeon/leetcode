package queue

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
	for this.ints[0] < t-3000{
		this.ints = this.ints[1:]
	}
	
	return  len(this.ints)
}
