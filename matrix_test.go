package lingo

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var valuePositionsForMatrix1 = [][]int{
	{0, 0, 1},
	{0, 1, 2},
	{0, 2, 3},
	{1, 0, 2},
	{1, 1, 4},
	{1, 2, 5},
}

var valuePositionsForMatrix1ByMatrix2 = [][]int{
	{0, 0, 19},
	{0, 1, 21},
	{1, 0, 34},
	{1, 1, 37},
}

var matrix1 = Matrix{
	{1, 2, 3},
	{2, 4, 5},
}

var matrix2 = Matrix{
	{1, 2},
	{3, 2},
	{4, 5},
}

func TestRows(t *testing.T) {
	assert.Equal(t, 2, matrix1.Rows(), "Number of rows should be 2")
}

func TestColumns(t *testing.T) {
	assert.Equal(t, 3, matrix1.Columns(), "Number of columns should be 3")
}

func TestDimensions(t *testing.T) {
	rows, columns := matrix1.Dimensions()
	assert.Equal(t, 2, rows, "Number of rows should be 2")
	assert.Equal(t, 3, columns, "Number of columns should be 3")
}

func TestValueAtPosition(t *testing.T) {
	for _, v := range valuePositionsForMatrix1 {
		val, err := matrix1.Value(v[0], v[1])
		assert.Equal(t, float64(v[2]), val, fmt.Sprintf("Value at position (%v, %v) should be %v", v[0], v[1], v[2]))
		assert.Equal(t, nil, err, "error should be nil")
	}
}

func TestValueAtPositionInvalidRow(t *testing.T) {
	val, err := matrix1.Value(3, 0)
	assert.Equal(t, float64(0), val, "value should be 0")
	assert.Equal(t, errors.New("value does not exist"), err, "error should be 'value does not exist'")
}

func TestValueAtPositionInvalidColumn(t *testing.T) {
	val, err := matrix1.Value(2, 4)
	assert.Equal(t, float64(0), val, "value should be 0")
	assert.Equal(t, errors.New("value does not exist"), err, "error should be 'value does not exist'")
}

func TestMultiply(t *testing.T) {
	m, err := matrix1.Multiply(matrix2)
	assert.Equal(t, nil, err, "error should be nil")

	for _, v := range valuePositionsForMatrix1ByMatrix2 {
		val, err := m.Value(v[0], v[1])
		assert.Equal(t, float64(v[2]), val, fmt.Sprintf("Value at position (%v, %v) should be %v", v[0], v[1], v[2]))
		assert.Equal(t, nil, err, "error should be nil")
	}
}
