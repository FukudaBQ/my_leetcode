package leetcode155

import (
	"encoding/json"
	"fmt"
	"math"
)

type MinStack struct {
	data     []int
	smallest int
	sIndex   []int
}

func Constructor() MinStack {
	return MinStack{
		data:     []int{},
		smallest: math.MaxInt32,
		sIndex:   []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if val < this.smallest {
		this.smallest = val
		this.sIndex = append(this.sIndex, len(this.data)-1)
	} else {
		if len(this.sIndex) > 0 {
			this.sIndex = append(this.sIndex, this.sIndex[len(this.sIndex)-1])
		} else {
			this.sIndex = append(this.sIndex, 0)
		}
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.sIndex = this.sIndex[:len(this.sIndex)-1]
	if len(this.data) > 0 {
		this.smallest = this.data[this.sIndex[len(this.sIndex)-1]]
	} else {
		this.smallest = math.MaxInt32
	}

}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.smallest
}

func OperateMinStack(ops []string, vals []string) []interface{} {
	var stack MinStack
	var res []interface{}

	for i, op := range ops {
		switch op {
		case "MinStack":
			stack = Constructor()
			fmt.Println(nil)
			res = append(res, nil)
		case "push":
			var nums []int
			json.Unmarshal([]byte(vals[i]), &nums)
			stack.Push(nums[0])
			fmt.Println(nil)
			res = append(res, nil)
		case "pop":
			stack.Pop()
			fmt.Println(nil)
			res = append(res, nil)
		case "top":
			val := stack.Top()
			fmt.Println(val)
			res = append(res, val)
		case "getMin":
			val := stack.GetMin()
			fmt.Println(val)
			res = append(res, val)
		}
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := OperateMinStack([]string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}, []string{"", "-2", "0", "-3", "", "", "", ""})
	print(a)
}
