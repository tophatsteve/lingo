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

var valuePositionsForMatrix1AddMatrix3 = [][]int{
	{0, 0, 4},
	{0, 1, 7},
	{0, 2, 4},
	{1, 0, 6},
	{1, 1, 8},
	{1, 2, 8},
}

var valuePositionsForMatrix3SubtractMatrix1 = [][]int{
	{0, 0, 2},
	{0, 1, 3},
	{0, 2, -2},
	{1, 0, 2},
	{1, 1, 0},
	{1, 2, -2},
}

var valuePositionsForMatrix1ScaledBy4 = [][]int{
	{0, 0, 4},
	{0, 1, 8},
	{0, 2, 12},
	{1, 0, 8},
	{1, 1, 16},
	{1, 2, 20},
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

var matrix3 = Matrix{
	{3, 5, 1},
	{4, 4, 3},
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

func TestMultiplyIncompatible(t *testing.T) {
	_, err := matrix1.Multiply(matrix3)
	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
}

func TestAdd(t *testing.T) {
	m, _ := matrix1.Add(matrix3)

	for _, v := range valuePositionsForMatrix1AddMatrix3 {
		val, err := m.Value(v[0], v[1])
		assert.Equal(t, float64(v[2]), val, fmt.Sprintf("Value at position (%v, %v) should be %v", v[0], v[1], v[2]))
		assert.Equal(t, nil, err, "error should be nil")
	}
}

func TestAddIncompatible(t *testing.T) {
	_, err := matrix1.Add(matrix2)
	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
}

func TestSubtract(t *testing.T) {
	m, _ := matrix3.Subtract(matrix1)

	for _, v := range valuePositionsForMatrix3SubtractMatrix1 {
		val, err := m.Value(v[0], v[1])
		assert.Equal(t, float64(v[2]), val, fmt.Sprintf("Value at position (%v, %v) should be %v", v[0], v[1], v[2]))
		assert.Equal(t, nil, err, "error should be nil")
	}
}

func TestSubtractIncompatible(t *testing.T) {
	_, err := matrix1.Subtract(matrix2)
	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
}

func TestScale(t *testing.T) {
	m := matrix1.Scale(4)

	for _, v := range valuePositionsForMatrix1ScaledBy4 {
		val, err := m.Value(v[0], v[1])
		assert.Equal(t, float64(v[2]), val, fmt.Sprintf("Value at position (%v, %v) should be %v", v[0], v[1], v[2]))
		assert.Equal(t, nil, err, "error should be nil")
	}
}
