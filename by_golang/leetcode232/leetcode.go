package leetcode232

type MyQueue struct {
	data1 stack
	data2 stack
}

func Constructor() MyQueue {
	data1 := stack{data: []int{}}
	data2 := stack{data: []int{}}
	return MyQueue{data1: data1, data2: data2}
}

func (this *MyQueue) Push(x int) {
	for this.data1.len() > 0 {
		this.data2.stackPush(this.data1.stackPop())
	}
	this.data1.stackPush(x)
	for this.data2.len() > 0 {
		this.data1.stackPush(this.data2.stackPop())
	}
}

func (this *MyQueue) Pop() int {
	return this.data1.stackPop()
}

func (this *MyQueue) Peek() int {
	return this.data1.stackPeek()
}

func (this *MyQueue) Empty() bool {
	return this.data1.len() == 0
}

type stack struct {
	data []int
}

func (this *stack) stackPush(x int) {
	this.data = append(this.data, x)
}

func (this *stack) stackPop() int {
	res := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return res
}

func (this *stack) stackPeek() int {
	res := this.data[len(this.data)-1]
	return res
}

func (this *stack) len() int {
	return len(this.data)
}
