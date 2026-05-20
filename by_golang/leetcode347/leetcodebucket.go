package leetcode347

func topKFrequentBucket(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	buckets := make([][]int, len(nums)+1)
	for v, f := range freq {
		buckets[f] = append(buckets[f], v)
	}
	res := make([]int, 0, k)
	for f := len(buckets) - 1; f >= 1 && len(res) < k; f-- {
		for _, v := range buckets[f] {
			res = append(res, v)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
