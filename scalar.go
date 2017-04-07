package lingo

import (
	"errors"
	"fmt"
)

// Scalar represents a single float64 number.
type Scalar float64

// Order returns the tensor order. For a scalar, this is always 0.
func (s Scalar) Order() int {
	return 0
}

// Rows returns the number of rows in the tensor. For a scalar, this is always 0.
func (s Scalar) Rows() int {
	return 0
}

// Columns returns the number of columns in the tensor. For a scalar, this is always 0.
func (s Scalar) Columns() int {
	return 0
}

func (s Scalar) Value(position ...int) (float64, error) {
	if len(position) == 0 {
		return float64(s), nil
	}

	if len(position) == 1 && position[0] == 0 {
		return float64(s), nil
	}

	if len(position) == 2 && position[0] == 0 && position[1] == 0 {
		return float64(s), nil
	}

	return 0, fmt.Errorf("no value at position %v", position)
}

func (s Scalar) SetValue(value float64, position ...int) (Tensor, error) {
	if len(position) == 0 {
		return Scalar(value), nil
	}

	return nil, fmt.Errorf("cannot set value at position %v", position)
}

func (s Scalar) String() string {
	return fmt.Sprintf("%v\n", float64(s))
}

// Reshape converts a scalar to either a Vector or element Matrix.
// The value of the scalar is broadcast to fill the dimensions of
// the new Tensor.
func (s Scalar) Reshape(dims ...int) (Tensor, error) {
	if len(dims) > 2 {
		return nil, errors.New("more than 2 dimensions is not supported")
	}

	if len(dims) == 0 {
		return s, nil
	}

	if len(dims) == 1 {
		v, err := s.Value()
		if err != nil {
			return nil, err
		}
		var vec = Vector{}
		for x := 0; x < dims[0]; x++ {
			vec = append(vec, v)
		}

		return vec, nil
	}

	if len(dims) == 2 {
		v, err := s.Value()
		if err != nil {
			return nil, err
		}
		var mat = Matrix{}
		for x := 0; x < dims[0]; x++ {
			var row []float64
			for y := 0; y < dims[1]; y++ {
				row = append(row, v)
			}
			mat = append(mat, row)
		}
		return mat, nil
	}

	return nil, errors.New("dimensions do not match")
}
