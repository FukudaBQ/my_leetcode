package leetcode146

type LRUCache struct {
	data     map[int]*linkedNode
	lru      *linkedNode
	mru      *linkedNode
	capacity int
}

type linkedNode struct {
	Val    int
	Key    int
	Next   *linkedNode
	Before *linkedNode
}

func Constructor(capacity int) LRUCache {
	data := make(map[int]*linkedNode, capacity)
	lru := &linkedNode{Val: -1, Key: -1, Next: nil, Before: nil}
	mru := &linkedNode{Val: -1, Key: -1, Next: nil, Before: lru}
	lru.Next = mru
	return LRUCache{data: data, capacity: capacity, lru: lru, mru: mru}
}

func (this *LRUCache) Get(key int) int {
	res, ok := this.data[key]
	if !ok {
		return -1
	}
	res.Before.Next = res.Next
	res.Next.Before = res.Before
	this.mru.Before.Next = res
	res.Before = this.mru.Before
	res.Next = this.mru
	this.mru.Before = res
	return res.Val
}

func (this *LRUCache) Put(key int, value int) {
	if res, ok := this.data[key]; ok {
		res.Val = value
		this.Get(key)
		return
	}

	if this.capacity == len(this.data) {
		lruNode := this.lru.Next
		delete(this.data, lruNode.Key)
		this.lru.Next = lruNode.Next
		lruNode.Next.Before = this.lru
	}

	res := &linkedNode{Key: key, Val: value}
	this.data[key] = res

	res.Before = this.mru.Before
	res.Next = this.mru
	this.mru.Before.Next = res
	this.mru.Before = res
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
