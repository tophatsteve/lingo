package lingo

import (
	"errors"
	"fmt"
)

// Matrix represents a 2 dimensional matrix of float64
type Matrix [][]float64

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
	if len(m)-1 < row {
		return []float64{}, errors.New("row does not exist")
	}
	return m[row], nil
}

// Column returns the specified column from the matrix
func (m Matrix) Column(column int) ([]float64, error) {
	if len(m)-1 < 0 {
		return []float64{}, errors.New("matrix has no rows")
	}

	if len(m[0])-1 < column {
		return []float64{}, errors.New("column does not exist")
	}

	r := []float64{}

	for x := range m {
		r = append(r, m[x][column])
	}

	return r, nil
}

// Dimensions returns the number of row and columns in the matrix
func (m Matrix) Dimensions() (int, int) {
	return m.Rows(), m.Columns()
}

// Add adds two matrices together
// An error is returned if the matrices are not compatible sizes.
func Add(m, o Matrix) (Matrix, error) {
	if !matchSize(m, o) {
		return nil, errors.New("incompatible matrices")
	}

	r := newZeroMatrix(m.Rows(), m.Columns())

	for mRows := 0; mRows < m.Rows(); mRows++ {
		for mCols := 0; mCols < m.Columns(); mCols++ {
			r[mRows][mCols] = m[mRows][mCols] + o[mRows][mCols]
		}
	}

	return r, nil
}

// Subtract subtracts one matrix from another.
// An error is returned if the matrices are not compatible sizes.
func Subtract(m, o Matrix) (Matrix, error) {
	if !matchSize(m, o) {
		return nil, errors.New("incompatible matrices")
	}

	r := newZeroMatrix(m.Rows(), m.Columns())

	for mRows := 0; mRows < m.Rows(); mRows++ {
		for mCols := 0; mCols < m.Columns(); mCols++ {
			r[mRows][mCols] = m[mRows][mCols] - o[mRows][mCols]
		}
	}

	return r, nil
}

// Multiply carries out element-wise multiplication of two matrices.
// An error is returned if the matrices are not compatible sizes.
func Multiply(m, o Matrix) (Matrix, error) {
	if !matchSize(m, o) {
		return nil, errors.New("incompatible matrices")
	}

	r := newZeroMatrix(m.Rows(), m.Columns())

	for mRows := 0; mRows < m.Rows(); mRows++ {
		for mCols := 0; mCols < m.Columns(); mCols++ {
			r[mRows][mCols] = m[mRows][mCols] * o[mRows][mCols]
		}
	}

	return r, nil
}

// Dot calculates the dot product of two matrices.
// An error is returned if the matrices are not compatible sizes.
func Dot(m, o Matrix) (Matrix, error) {
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

// Scale multiplies the elements of a matrix by a scalar value.
func Scale(m Matrix, scalar float64) Matrix {
	r := newZeroMatrix(m.Rows(), m.Columns())

	for mRows := 0; mRows < m.Rows(); mRows++ {
		for mCols := 0; mCols < m.Columns(); mCols++ {
			r[mRows][mCols] = m[mRows][mCols] * scalar
		}
	}

	return r
}

// Transpose transposes a matrix, converting its rows to columns
// and its columns to rows.
func Transpose(m Matrix) Matrix {
	r := newZeroMatrix(m.Columns(), m.Rows())

	for mRows := 0; mRows < m.Rows(); mRows++ {
		for mCols := 0; mCols < m.Columns(); mCols++ {
			r[mCols][mRows] = m[mRows][mCols]
		}
	}

	return r
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
