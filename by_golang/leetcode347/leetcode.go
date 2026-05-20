package leetcode347

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	topK := make([]int, 0, k)
	lowest := 0
	lowestIdx := 0
	for key, v := range freq {
		if len(topK) < k {
			topK = append(topK, key)
			if len(topK) == 1 || v < lowest {
				lowest = v
				lowestIdx = len(topK) - 1
			}
		} else if v > lowest {
			topK[lowestIdx] = key
			lowest = freq[topK[0]]
			lowestIdx = 0
			for i := 1; i < k; i++ {
				if f := freq[topK[i]]; f < lowest {
					lowest = f
					lowestIdx = i
				}
			}
		}
	}
	return topK
}

func Leetcode() {
	nums, k := []int{1, 1, 1, 2, 2, 3}, 2
	fmt.Printf("scan:   %v\n", topKFrequent(nums, k))
	fmt.Printf("heap:   %v\n", topKFrequentHeap(nums, k))
	fmt.Printf("bucket: %v\n", topKFrequentBucket(nums, k))
}
