package lingo

import (
	"errors"
	"fmt"
)

// Matrix represents a 2 dimensional matrix of float64
type Matrix [][]float64

func (m Matrix) Add(o Matrix) (Matrix, error) {
	if !matchSize(m, o) {
		return nil, errors.New("incompatible matrices")
	}

	return nil, nil
}

func (m Matrix) Subtract(o Matrix) (Matrix, error) {
	if !matchSize(m, o) {
		return nil, errors.New("incompatible matrices")
	}

	return nil, nil
}

func (m Matrix) Multiply(o Matrix) (Matrix, error) {
	if m.Columns() != o.Rows() {
		return nil, errors.New("incompatible matrices")
	}

	r := newZeroMatrix(m.Rows(), o.Columns())

	for mRows := 0; mRows < m.Rows(); mRows++ {
		for oCols := 0; oCols < o.Columns(); oCols++ {
			r[mRows][oCols] = 0
			for mCols := 0; mCols < m.Columns(); mCols++ {
				r[mRows][oCols] += m[mRows][mCols] * o[mCols][oCols]
			}
		}
	}

	return r, nil
}

// Scale multiplies a matrix by a scalar value
func (m Matrix) Scale(scalar float64) (Matrix, error) {
	return nil, nil
}

// Print writes the matrix to stdout
func (m Matrix) Print() {
	// print the matrix
	for _, row := range m {
		fmt.Printf("%v\n", row)
	}
}

// Rows returns the number of rows the matrix contains
func (m Matrix) Rows() int {
	return len(m)
}

// Columns returns the number of columns the matrix contains
func (m Matrix) Columns() int {
	if len(m) == 0 {
		return 0
	}

	return len(m[0])
}

// Value returns the value at a specific, zero-based position in the matrix
func (m Matrix) Value(row, column int) (float64, error) {
	if len(m)-1 < row {
		return 0, errors.New("value does not exist")
	}

	if len(m[row])-1 < column {
		return 0, errors.New("value does not exist")
	}

	return m[row][column], nil
}

// Row returns the specified row from the matrix
func (m Matrix) Row(row int) ([]float64, error) {
	return []float64{}, nil
}

// Column returns the specified column from the matrix
func (m Matrix) Column(column int) ([]float64, error) {
	return []float64{}, nil
}

// Dimensions returns the number of row and columns in the matrix
func (m Matrix) Dimensions() (int, int) {
	return m.Rows(), m.Columns()
}

func newZeroMatrix(rows, columns int) Matrix {
	m := Matrix{}
	for x := 0; x < rows; x++ {
		r := make([]float64, columns, columns)
		m = append(m, r)
	}
	return m
}

func matchSize(m1, m2 Matrix) bool {
	if m1.Columns() != m2.Columns() || m1.Rows() != m2.Rows() {
		return false
	}
	return true
}
