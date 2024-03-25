package hashset

type MyHashSet struct {
	set [1000000/8 + 1]byte
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	byt := key / 8
	bit := key % 8
	this.set[byt] |= 1 << bit
}

func (this *MyHashSet) Remove(key int) {
	byt := key / 8
	bit := key % 8
	this.set[byt] &= ^(1 << bit)
}

func (this *MyHashSet) Contains(key int) bool {
	byt := key / 8
	bit := key % 8
	val := byte(1) << bit
	return this.set[byt]&val == val
}

type MyHashSetImpl0 struct {
	x map[int]struct{}
}

// func Constructor() MyHashSetImpl0 {
// 	return MyHashSetImpl0{x: map[int]struct{}{}}
// }

func (this *MyHashSetImpl0) Add(key int) {
	this.x[key] = struct{}{}
}

func (this *MyHashSetImpl0) Remove(key int) {
	delete(this.x, key)
}

func (this *MyHashSetImpl0) Contains(key int) bool {
	_, ok := this.x[key]
	return ok
}
