package hash_table

type LinkNode struct {
	value int
	next  *LinkNode
}

// 706 
type MyHashMap struct {
	list [10000]LinkNode
}

/** Initialize your data structure here. */
func ConstructorMap() MyHashMap {
	return MyHashMap{
		list: [10000]LinkNode{},
	}
}

func (this *MyHashMap) hashKey(key int) int {
	M := 101
	return key % M
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {

	hk := this.hashKey(key)
	node := &this.list[hk]
	if node.value == 0 {
		node.value = value
	} else {
		valNode := LinkNode{
			value: value,
		}
		node.next = &valNode
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hk := this.hashKey(key)
	node := this.list[hk]
	return node.value
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
