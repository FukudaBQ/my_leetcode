package leetcode692

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	freq := make(map[string]int, k)
	for _, word := range words {
		freq[word]++
	}
	bucket := make(map[int][]string, len(words))
	for word, f := range freq {
		var iBucket []string
		if _, ok := bucket[f]; !ok {
			iBucket = make([]string, 0, k)
		} else {
			iBucket = bucket[f]
		}
		bucket[f] = append(iBucket, word)
	}
	res := make([]string, 0, k)
	for i := len(words); i > 0 && len(res) < k; i-- {
		if _, ok := bucket[i]; ok {
			sort.Strings(bucket[i])
			for _, word := range bucket[i] {
				res = append(res, word)
			}
		}
	}
	return res[:k]
}

func Leetcode() {
	nums, k := []string{"i", "love", "leetcode", "i", "love", "coding"}, 1
	fmt.Printf("scan:   %v\n", topKFrequent(nums, k))
}
