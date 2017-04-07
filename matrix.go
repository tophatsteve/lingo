package lingo

import (
	"errors"
	"fmt"
)

// Matrix represents a 2 dimensional matrix of float64.
type Matrix [][]float64

// Order returns the tensor order. For a matrix, this is always 2.
func (m Matrix) Order() int {
	return 2
}

// Rows returns the number of rows the matrix contains.
func (m Matrix) Rows() int {
	return len(m)
}

// Columns returns the number of columns the matrix contains.
func (m Matrix) Columns() int {
	if len(m) == 0 {
		return 0
	}

	return len(m[0])
}

// Value returns the value at a specific, zero-based position in the matrix.
func (m Matrix) Value(position ...int) (float64, error) {
	if len(position) != 2 {
		return 0, errors.New("matrix positions are 2 dimensional")
	}
	if len(m)-1 < position[0] {
		return 0, errors.New("value does not exist")
	}

	if len(m[position[0]])-1 < position[1] {
		return 0, errors.New("value does not exist")
	}

	return m[position[0]][position[1]], nil
}

// SetValue sets the value at a position in the matrix.
func (m Matrix) SetValue(value float64, position ...int) (Tensor, error) {
	if len(position) != 2 {
		return nil, errors.New("matrix positions are 2 dimensional")
	}
	if position[0] >= m.Rows() || position[1] > m.Columns() {
		return nil, errors.New("position is not in Matrix")
	}

	m[position[0]][position[1]] = value

	return m, nil
}

// String returns a string representation of the matrix.
func (m Matrix) String() string {
	output := ""

	for _, v := range m {
		output += fmt.Sprintf("%v\n", v)
	}

	return output
}

// Reshape converts the matrix m into a new matrix with dimensions r,c.
func (m Matrix) Reshape(dims ...int) (Tensor, error) {
	// check matching size
	if len(dims) == 0 {
		// scalar
		if m.Rows() != 1 || m.Columns() != 1 {
			return nil, errors.New("can only reshape a 1x1 matrix to a scalar")
		}
	}

	if len(dims) == 1 {
		// vector
		if dims[0] != m.Rows()*m.Columns() {
			return nil, errors.New("dimensions do not match")
		}
	}

	if len(dims) == 2 {
		// matrix
		if dims[0]*dims[1] != m.Rows()*m.Columns() {
			return nil, errors.New("dimensions do not match")
		}
	}

	return nil, errors.New("more than 2 dimensions is not supported")
}
