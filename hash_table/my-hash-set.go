package hash_table

// 705

type MyHashSet struct {
	set map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {

	return MyHashSet{set: map[int]int{}}
}


func (this *MyHashSet) Add(key int)  {
	this.set[key] ++
}


func (this *MyHashSet) Remove(key int)  {
	delete(this.set, key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if this.set[key] > 0 {
		return true
	}

	return false
}