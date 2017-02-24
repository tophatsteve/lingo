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

func TestColumnsEmptyMatrix(t *testing.T) {
	m := Matrix{}
	assert.Equal(t, 0, m.Columns(), "Number of columns should be 0")
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
	val, err := matrix1.Value(1, 4)
	assert.Equal(t, float64(0), val, "value should be 0")
	assert.Equal(t, errors.New("value does not exist"), err, "error should be 'value does not exist'")
}

func TestRow(t *testing.T) {
	e := []float64{2, 4, 5}
	row, err := matrix1.Row(1)

	assert.Equal(t, len(e), len(row), "row is not the correct length")

	for i, v := range e {
		assert.Equal(t, v, row[i], fmt.Sprintf("value '%v' should be '%v'", row[i], v))
	}

	assert.Equal(t, nil, err, "error should be nil")
}

func TestInvalidRow(t *testing.T) {
	row, err := matrix1.Row(2)
	assert.Equal(t, 0, len(row), "row length should be 0")
	assert.Equal(t, errors.New("row does not exist"), err, "error should be 'row does not exist'")
}

func TestColumn(t *testing.T) {
	e := []float64{3, 5}
	col, err := matrix1.Column(2)

	assert.Equal(t, len(e), len(col), "column is not the correct length")

	for i, v := range e {
		assert.Equal(t, v, col[i], fmt.Sprintf("value '%v' should be '%v'", col[i], v))
	}

	assert.Equal(t, nil, err, "error should be nil")
}

func TestInvalidColumn(t *testing.T) {
	col, err := matrix1.Column(3)
	assert.Equal(t, 0, len(col), "column length should be 0")
	assert.Equal(t, errors.New("column does not exist"), err, "error should be 'column does not exist'")
}

func TestColumnNoRows(t *testing.T) {
	m := Matrix{}
	col, err := m.Column(0)
	assert.Equal(t, 0, len(col), "column length should be 0")
	assert.Equal(t, errors.New("matrix has no rows"), err, "error should be 'matrix has no rows'")
}

func TestDot(t *testing.T) {
	m, err := Dot(matrix1, matrix2)
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

func TestDotIncompatible(t *testing.T) {
	_, err := Dot(matrix1, matrix3)
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

func TestMultiply(t *testing.T) {
	m, _ := Multiply(matrix3, matrix1)
	e := Matrix{
		{3, 10, 3},
		{8, 16, 15},
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
	_, err := Multiply(matrix1, matrix2)
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
