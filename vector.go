package lingo

import (
	"errors"
	"fmt"
)

type Vector []float64

// Order returns the tensor order. For a vector, this is always 1.
func (v Vector) Order() int {
	return 1
}

func (v Vector) Rows() int {
	if len(v) == 0 {
		return 0
	}

	return 1
}

func (v Vector) Columns() int {
	return len(v)
}

func (v Vector) Value(position ...int) (float64, error) {
	if len(position) == 1 && len(v) > position[0] {
		return float64(v[position[0]]), nil
	}

	if len(position) == 2 && position[0] == 0 && len(v) > position[1] {
		return float64(v[position[1]]), nil
	}

	return 0, fmt.Errorf("no value at position %v", position)
}

func (v Vector) SetValue(value float64, position ...int) (Tensor, error) {
	return nil, fmt.Errorf("cannot set value at position %v", position)
}

func (v Vector) String() string {
	return fmt.Sprintf("%v\n", Scalar(v[0]))
}

// Reshape converts a vector to a r,c dimension matrix.
func (v Vector) Reshape(dims ...int) (Tensor, error) {
	// check matching size
	if len(dims) == 0 {
		// scalar
		if len(v) > 1 {
			return nil, errors.New("dimensions do not match")
		}
		return Scalar(v[0]), nil
	}

	if len(dims) == 1 {
		// vector
		if dims[0] != len(v) {
			return nil, errors.New("dimensions do not match")
		}
		return v, nil
	}

	if len(dims) == 2 {
		// matrix
		if dims[0]*dims[1] != len(v) {
			return nil, errors.New("dimensions do not match")
		}
		m := Matrix{}
		for x := 0; x < dims[0]; x++ {
			var row []float64
			for y := 0; y < dims[1]; y++ {
				val, err := v.Value((x * dims[1]) + y)
				if err != nil {
					return nil, err
				}
				row = append(row, val)
			}
			m = append(m, row)
		}
		return m, nil
	}

	return nil, errors.New("more than 2 dimensions is not supported")
}
