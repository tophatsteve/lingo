package lingo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScalarRows(t *testing.T) {
	s := Scalar(1)
	assert.Equal(t, 0, s.Rows(), "scalar rows not equal to 0")
}

func TestScalarColumns(t *testing.T) {
	s := Scalar(1)
	assert.Equal(t, 0, s.Columns(), "scalar columns not equal to 0")
}

func TestScalarValue(t *testing.T) {
	s := Scalar(1)
	v, err := s.Value()
	assert.Equal(t, nil, err, "err is not nil")
	assert.Equal(t, float64(1), v, "value is not correct")
}

func TestScalarValueAtPosition1Dimension(t *testing.T) {
	s := Scalar(1)
	v, err := s.Value(0)
	assert.Equal(t, nil, err, "err is not nil")
	assert.Equal(t, float64(1), v, "value is not correct")
}

func TestScalarValueAtPosition2Dimension(t *testing.T) {
	s := Scalar(1)
	v, err := s.Value(0, 0)
	assert.Equal(t, nil, err, "err is not nil")
	assert.Equal(t, float64(1), v, "value is not correct")
}

func TestScalarValueAtPositionInvalidDimensionValue(t *testing.T) {
	s := Scalar(1)
	_, err := s.Value(1)
	assert.Equal(t, errors.New("no value at position [1]"), err, "err is not as expected")
}

func TestScalarValueAtPositionInvalidNumberOfDimensions(t *testing.T) {
	s := Scalar(1)
	_, err := s.Value(0, 0, 0)
	assert.Equal(t, errors.New("no value at position [0 0 0]"), err, "err is not as expected")
}

func TestScalarSetValue(t *testing.T) {
	s := Scalar(1)
	s2, err := s.SetValue(2)
	assert.Equal(t, nil, err, "err is not nil")
	v, err := s2.Value()
	assert.Equal(t, nil, err, "err is not nil")
	assert.Equal(t, float64(2), v, "value is not correct")
}

func TestScalarSetValueAtPosition(t *testing.T) {
	s := Scalar(1)
	_, err := s.SetValue(2, 1)
	assert.Equal(t, errors.New("cannot set value at position [1]"), err, "err is not as expected")
}

func TestScalarReshapeToScalar(t *testing.T) {
	s := Scalar(1)
	s2, _ := s.Reshape()
	v, _ := s2.Value()
	assert.Equal(t, 0, s2.Order(), "reshaping scalar to scalar did not return a scalar")
	assert.Equal(t, float64(1), v, "reshaping scalar to scalar did not retain correct value")
}

func TestScalarReshapeToVector(t *testing.T) {
	s := Scalar(1)
	v, _ := s.Reshape(1)
	val, _ := v.Value(0)
	assert.Equal(t, 1, v.Order(), "reshaping scalar to vector did not return a vector")
	assert.Equal(t, float64(1), val, "reshaping scalar to vector did not retain correct value")
}

func TestScalarReshapeAndBroadcastToVector(t *testing.T) {
	s := Scalar(1)
	vectorLength := 4
	v, _ := s.Reshape(vectorLength)
	assert.Equal(t, 1, v.Order(), "reshaping scalar to vector did not return a vector")

	for x := 0; x < vectorLength; x++ {
		val, _ := v.Value(x)
		assert.Equal(t, float64(1), val, "broadcasting scalar to vector did not retain correct value")
	}
}

func TestScalarReshapeToMatrix(t *testing.T) {
	s := Scalar(1)
	m, _ := s.Reshape(1, 1)
	val, _ := m.Value(0, 0)
	assert.Equal(t, 2, m.Order(), "reshaping scalar to matrix did not return a matrix")
	assert.Equal(t, float64(1), val, "reshaping scalar to matrix did not retain correct value")
}

func TestScalarReshapeAndBroadcastToMatrix(t *testing.T) {
	s := Scalar(1)
	matrixRows := 4
	matrixColumns := 3
	v, _ := s.Reshape(matrixRows, matrixColumns)
	assert.Equal(t, 2, v.Order(), "reshaping scalar to matrix did not return a matrix")

	for x := 0; x < matrixRows; x++ {
		for y := 0; y < matrixColumns; y++ {
			val, _ := v.Value(x, y)
			assert.Equal(t, float64(1), val, "broadcasting scalar to matrix did not retain correct value")
		}
	}
}
