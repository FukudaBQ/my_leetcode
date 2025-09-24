package leetcode641

type MyCircularDeque struct {
	length    int
	maxLength int
	data      []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		length:    0,
		maxLength: k,
		data:      make([]int, 0, k),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.length == this.maxLength {
		return false
	}
	this.length++
	this.data = append([]int{value}, this.data...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.length == this.maxLength {
		return false
	}
	this.length++
	this.data = append(this.data, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.length == 0 {
		return false
	}
	this.data = this.data[1:this.length]
	this.length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.length == 0 {
		return false
	}
	this.length--
	this.data = this.data[0:this.length]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.length == 0 {
		return -1
	}
	return this.data[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.length == 0 {
		return -1
	}
	return this.data[this.length-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.maxLength
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
