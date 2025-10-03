package leetcode304

import "fmt"

type NumMatrix struct {
	data      [][]int
	prefixSum [][]int
	height    int
	length    int
}

func Constructor(matrix [][]int) NumMatrix {
	h := len(matrix)
	if h == 0 {
		return NumMatrix{}
	}
	w := len(matrix[0])
	prefixSum := make([][]int, h+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, w+1)
	}

	for i := 1; i <= h; i++ {
		realRow := h - i
		rowSum := 0
		for j := 1; j <= w; j++ {
			rowSum += matrix[realRow][j-1]
			prefixSum[i][j] = prefixSum[i-1][j] + rowSum
		}
	}

	return NumMatrix{
		data:      matrix,
		prefixSum: prefixSum,
		height:    h,
		length:    w,
	}
}

func (this *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	r1 := this.height - row1 - 1
	r2 := this.height - row2 - 1
	if r1 < r2 {
		r1, r2 = r2, r1
	}
	cLeft, cRight := col1, col2
	if cLeft > cRight {
		cLeft, cRight = cRight, cLeft
	}

	p := this.prefixSum
	return p[r2][cLeft] + p[r1+1][cRight+1] - p[r2][cRight+1] - p[r1+1][cLeft]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func Leetcode() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)
	fmt.Println(obj.SumRegion(2, 1, 4, 3)) // 8
	fmt.Println(obj.SumRegion(1, 1, 2, 2)) // 11
	fmt.Println(obj.SumRegion(1, 2, 2, 4)) // 12
}
