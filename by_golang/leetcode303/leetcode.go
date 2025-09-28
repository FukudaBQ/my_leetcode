package leetcode303

type NumArray struct {
	nums []int
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	sums[0] = nums[0]
	for i, j := range nums {
		sums[i+1] = sums[i] + j
	}
	return NumArray{nums: nums, sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
