package leetcode81

import "fmt"

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		//fmt.Printf("nums[mid] = %d, mid = %d, left = %d, right = %d\n", nums[mid], mid, l, r)
		if target == nums[mid] {
			return true
		}
		if r-l <= 2 {
			//fmt.Printf("nums[mid] = %d\n", nums[mid])
			if nums[l] == target {
				return true
			}
			if nums[mid] == target {
				return true
			}
			if nums[r] == target {
				return true
			}
			return false
		}
		if nums[mid] == nums[l] {
			l++
			continue
		}
		if nums[mid] > nums[l] {
			//fmt.Printf("111:nums[mid] = %d\n", nums[mid])
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			//fmt.Printf("222:nums[mid] = %d\n", nums[mid])
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		//fmt.Printf("END: nums[mid] = %d, mid = %d, left = %d, right = %d\n", nums[mid], mid, l, r)
	}
	return false
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2)
	fmt.Printf("is anagram: %v", a)
}
