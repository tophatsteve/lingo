package lingo

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// define test matrices
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
	e := [][]int{
		{0, 0, 1},
		{0, 1, 2},
		{0, 2, 3},
		{1, 0, 2},
		{1, 1, 4},
		{1, 2, 5},
	}

	for _, v := range e {
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
	m, err := Multiply(matrix1, matrix2)
	assert.Equal(t, nil, err, "error should be nil")

	e := Matrix{
		{19, 21},
		{34, 37},
	}

	for row := range e {
		for col := range e[row] {
			val, err := m.Value(row, col)
			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
			assert.Equal(t, nil, err, "error should be nil")
		}
	}
}

func TestMultiplyIncompatible(t *testing.T) {
	_, err := Multiply(matrix1, matrix3)
	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
}

func TestAdd(t *testing.T) {
	m, _ := Add(matrix1, matrix3)
	e := Matrix{
		{4, 7, 4},
		{6, 8, 8},
	}

	for row := range e {
		for col := range e[row] {
			val, err := m.Value(row, col)
			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
			assert.Equal(t, nil, err, "error should be nil")
		}
	}
}

func TestAddIncompatible(t *testing.T) {
	_, err := Add(matrix1, matrix2)
	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
}

func TestSubtract(t *testing.T) {
	m, _ := Subtract(matrix3, matrix1)
	e := Matrix{
		{2, 3, -2},
		{2, 0, -2},
	}

	for row := range e {
		for col := range e[row] {
			val, err := m.Value(row, col)
			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
			assert.Equal(t, nil, err, "error should be nil")
		}
	}
}

func TestSubtractIncompatible(t *testing.T) {
	_, err := Subtract(matrix1, matrix2)
	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
}

func TestScale(t *testing.T) {
	m := Scale(matrix1, 4)
	e := Matrix{
		{4, 8, 12},
		{8, 16, 20},
	}

	for row := range e {
		for col := range e[row] {
			val, err := m.Value(row, col)
			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
			assert.Equal(t, nil, err, "error should be nil")
		}
	}
}

func TestTranspose(t *testing.T) {
	m := Transpose(matrix1)
	e := Matrix{
		{1, 2},
		{2, 4},
		{3, 5},
	}

	for row := range e {
		for col := range e[row] {
			val, err := m.Value(row, col)
			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
			assert.Equal(t, nil, err, "error should be nil")
		}
	}
}
