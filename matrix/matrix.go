package matrix

import (
	"fmt"
	"strings"
)

// Matrix represents a matrix of arbitrary size.
// This is also the same type used to represent vectors.
type Matrix struct {
	data [][]int
	w    int
	h    int
}

// New creates a new Matrix based off a nested array.
func New(init [][]int) Matrix {
	h := len(init)
	w := 0
	if h > 0 {
		w = len(init[0])
	}

	return Matrix{
		data: init,
		w:    w,
		h:    h,
	}
}

// NewVector is syntax sugar for creating matrices with a height of one.
func NewVector(init []int) Matrix {
	return New([][]int{init})
}

func (m Matrix) String() string {
	row := strings.Repeat("%v ", m.w)
	entirety := strings.Repeat(row[:len(row)-1]+"\n", m.h)
	held := make([]interface{}, m.h*m.w)
	for y, row := range m.data {
		for x, item := range row {
			held[y*m.w+x] = item
		}
	}
	return fmt.Sprintf(entirety, held...)
}

// Copy performs an explicit copy of a Matrix.
func (m Matrix) Copy() Matrix {
	data := make([][]int, m.h)
	for i := range data {
		data[i] = make([]int, m.w)
		copy(data[i], m.data[i])
	}
	return Matrix{
		data: data,
		w:    m.w,
		h:    m.h,
	}
}

// AddMe performs in-place addition of two matrices.
// This is made available when performance is needed, to avoid unnecessary copies.
// The modified Matrix is returned to allow chaining.
func (m Matrix) AddMe(to Matrix) Matrix {
	for y, row := range m.data {
		for x := range row {
			row[x] += to.data[y][x]
		}
	}
	return m
}

// Add performs a copying addition of two matrices.
func (m Matrix) Add(to Matrix) Matrix {
	return m.Copy().AddMe(to)
}

// SubtractMe performs in-place subtraction of two matrices.
// See AddMe for notes on why this is made available.
func (m Matrix) SubtractMe(to Matrix) Matrix {
	for y, row := range m.data {
		for x := range row {
			row[x] -= to.data[y][x]
		}
	}
	return m
}

// Subtract performs a copying subtraction of two matrices.
func (m Matrix) Subtract(to Matrix) Matrix {
	return m.Copy().SubtractMe(to)
}

// DotProduct does a copying dot product between a vector and a matrix.
func (v Matrix) DotProduct(m Matrix) Matrix {
	data := make([][]int, m.w)
	for i := range data {
		data[i] = v.data[0]
	}
	return New(data).Multiply(m)
}

// Multiply performs a multiplication between two matrices.
func (m Matrix) Multiply(to Matrix) Matrix {
	data := make([][]int, m.h)
	for i := range data {
		data[i] = make([]int, m.h)
	}
	ret := New(data)

	for leftRowIndex, leftRow := range m.data {
		for rightColumnIndex := 0; rightColumnIndex < m.h; rightColumnIndex++ {
			for depth := 0; depth < m.w; depth++ {
				ret.data[leftRowIndex][rightColumnIndex] += leftRow[depth] * to.data[depth][rightColumnIndex]
			}
		}
	}

	return ret
}

// Transpose me performs an in-place transposition of a matrix.
// See AddMe for notes on why this is made available.
// Currently this may seem silly when you look at the implimentation;
// But once a single slice is used instead of nested ones, this can be done
// without any allocation.
func (m Matrix) TransposeMe() Matrix {
	data := make([][]int, m.w)
	for y := range data {
		data[y] = make([]int, m.h)
		for x := range data[y] {
			data[y][x] = m.data[x][y]
		}
	}

	m.data = data
	m.w, m.h = m.h, m.w

	return m
}

// Transpose does a copying Transpose of a Matrix.
func (m Matrix) Transpose() Matrix {
	return m.Copy().TransposeMe()
}
