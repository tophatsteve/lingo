// Package lingo provides matrix and vector operations
package lingo

import (
	"errors"
)

type Tensor interface {
	Order() int
	Rows() int
	Columns() int
	Value(position ...int) (float64, error)
	SetValue(value float64, position ...int) (Tensor, error)
	Reshape(dims ...int) (Tensor, error)
	String() string
}

// Add adds two compatible tensors together.
// An error is returned if the tensors are not compatible sizes.
func Add(m, o Tensor) (Tensor, error) {
	if !matchSize(m, o) {
		return nil, errors.New("incompatible matrices")
	}

	// Scalar
	if m.Order() == 0 {
		mValue, err := m.Value()
		if err != nil {
			return nil, err
		}
		oValue, err := o.Value()
		if err != nil {
			return nil, err
		}
		return Scalar(float64(mValue) + float64(oValue)), nil
	}

	var r Tensor
	var err error

	if m.Rows() == 1 {
		r, err = newZeroTensor(m.Columns())
		if err != nil {
			return nil, err
		}
	}

	if m.Rows() > 2 {
		r, err = newZeroTensor(m.Rows(), m.Columns())
		if err != nil {
			return nil, err
		}
	}

	for mRows := 0; mRows < m.Rows(); mRows++ {
		for mCols := 0; mCols < m.Columns(); mCols++ {
			mValue, _ := m.Value(mRows, mCols)
			oValue, _ := o.Value(mRows, mCols)
			r.SetValue(mValue+oValue, mRows, mCols)
		}
	}

	return nil, nil
}

func matchSize(m1, m2 Tensor) bool {
	if m1.Columns() != m2.Columns() || m1.Rows() != m2.Rows() {
		return false
	}
	return true
}

func newZeroTensor(dims ...int) (Tensor, error) {
	if len(dims) == 0 {
		s := Scalar(0)
		return s, nil
	}

	if len(dims) == 1 {
		v := Vector(make([]float64, dims[0]))
		return v, nil
	}

	if len(dims) == 2 {
		m := Matrix{}
		for x := 0; x < dims[0]; x++ {
			r := make([]float64, dims[1], dims[1])
			m = append(m, r)
		}
		return m, nil
	}

	return nil, errors.New("more than 2 dimensions is not supported")
}
