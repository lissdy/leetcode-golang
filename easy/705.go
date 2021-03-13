package main

type MyHashSet struct {
	set []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{set: make([]int, 1000001)}
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.set[key] = 0
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if this.set[key] == 1 {
		return true
	}
	return false
}
