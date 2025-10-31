package leetcode33

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		//fmt.Printf("nums[mid] = %d, mid = %d, left = %d, right = %d\n", nums[mid], mid, l, r)
		if target == nums[mid] {
			return mid
		}
		if r-l <= 2 {
			//fmt.Printf("nums[mid] = %d\n", nums[mid])
			if nums[l] == target {
				return l
			}
			if nums[mid] == target {
				return mid
			}
			if nums[r] == target {
				return r
			}
			return -1
		}
		if nums[mid] > nums[l] {
			//fmt.Printf("111:nums[mid] = %d\n", nums[mid])
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//fmt.Printf("222:nums[mid] = %d\n", nums[mid])
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		//fmt.Printf("END: nums[mid] = %d, mid = %d, left = %d, right = %d\n", nums[mid], mid, l, r)
	}
	return -1
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := search([]int{4, 5, 6, 7, 0, 1, 2}, 8)
	fmt.Printf("is anagram: %v", a)
}
