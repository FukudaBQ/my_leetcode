package leetcode433

import "fmt"

func diffByOne(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return diff == 1
}

func minMutation(startGene string, endGene string, bank []string) int {
	bankSet := make(map[string]struct{}, len(bank))
	for _, g := range bank {
		bankSet[g] = struct{}{}
	}
	if _, ok := bankSet[endGene]; !ok {
		return -1
	}

	visited := make(map[string]struct{})
	visited[startGene] = struct{}{}
	current := []string{startGene}
	steps := 0

	for len(current) > 0 {
		next := make([]string, 0)
		for _, gene := range current {
			if gene == endGene {
				return steps
			}
			for candidate := range bankSet {
				if _, seen := visited[candidate]; seen {
					continue
				}
				if diffByOne(gene, candidate) {
					visited[candidate] = struct{}{}
					next = append(next, candidate)
				}
			}
		}
		current = next
		steps++
	}

	return -1
}

func Leetcode() {
	a := minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"})
	fmt.Printf("min mutation: %v\n", a)
}
