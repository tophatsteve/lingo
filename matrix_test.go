package lingo

// import (
// 	"errors"
// 	"fmt"
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// )

// // define test matrices
// var matrix1 = Matrix{
// 	{1, 2, 3},
// 	{2, 4, 5},
// }

// var matrix2 = Matrix{
// 	{1, 2},
// 	{3, 2},
// 	{4, 5},
// }

// var matrix3 = Matrix{
// 	{3, 5, 1},
// 	{4, 4, 3},
// }

// func TestRows(t *testing.T) {
// 	assert.Equal(t, 2, matrix1.Rows(), "Number of rows should be 2")
// }

// func TestColumns(t *testing.T) {
// 	assert.Equal(t, 3, matrix1.Columns(), "Number of columns should be 3")
// }

// func TestColumnsEmptyMatrix(t *testing.T) {
// 	m := Matrix{}
// 	assert.Equal(t, 0, m.Columns(), "Number of columns should be 0")
// }

// func TestValueAtPosition(t *testing.T) {
// 	e := [][]int{
// 		{0, 0, 1},
// 		{0, 1, 2},
// 		{0, 2, 3},
// 		{1, 0, 2},
// 		{1, 1, 4},
// 		{1, 2, 5},
// 	}

// 	for _, v := range e {
// 		val, err := matrix1.Value(v[0], v[1])
// 		assert.Equal(t, float64(v[2]), val, fmt.Sprintf("Value at position (%v, %v) should be %v", v[0], v[1], v[2]))
// 		assert.Equal(t, nil, err, "error should be nil")
// 	}
// }

// func TestValueAtPositionInvalidRow(t *testing.T) {
// 	val, err := matrix1.Value(3, 0)
// 	assert.Equal(t, float64(0), val, "value should be 0")
// 	assert.Equal(t, errors.New("value does not exist"), err, "error should be 'value does not exist'")
// }

// func TestValueAtPositionInvalidColumn(t *testing.T) {
// 	val, err := matrix1.Value(1, 4)
// 	assert.Equal(t, float64(0), val, "value should be 0")
// 	assert.Equal(t, errors.New("value does not exist"), err, "error should be 'value does not exist'")
// }

// func TestSetValue(t *testing.T) {
// 	assert.Equal(t, true, false, "TestSetValue not implemented")
// }

// func TestEqual(t *testing.T) {
// 	m1 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}
// 	m2 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}

// 	assert.Equal(t, true, m1.Equal(m2), "matrices should be equal")
// }

// func TestNotEqualRows(t *testing.T) {
// 	m1 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}
// 	m2 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 		{6, 7, 8},
// 	}

// 	assert.Equal(t, false, m1.Equal(m2), "matrices should not be equal")
// }

// func TestNotEqualColumns(t *testing.T) {
// 	m1 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}
// 	m2 := Matrix{
// 		{1, 2},
// 		{2, 4},
// 	}

// 	assert.Equal(t, false, m1.Equal(m2), "matrices should not be equal")
// }

// func TestNotEqualValues(t *testing.T) {
// 	m1 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}
// 	m2 := Matrix{
// 		{1, 2, 3},
// 		{2, 6, 5},
// 	}

// 	assert.Equal(t, false, m1.Equal(m2), "matrices should not be equal")
// }

// func TestDot(t *testing.T) {
// 	m, err := matrix1.Dot(matrix2)
// 	assert.Equal(t, nil, err, "error should be nil")

// 	e := Matrix{
// 		{19, 21},
// 		{34, 37},
// 	}

// 	for row := range e {
// 		for col := range e[row] {
// 			val, err := m.Value(row, col)
// 			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
// 			assert.Equal(t, nil, err, "error should be nil")
// 		}
// 	}
// }

// func TestDotIncompatible(t *testing.T) {
// 	_, err := matrix1.Dot(matrix3)
// 	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
// }

// func TestAdd(t *testing.T) {
// 	m, _ := matrix1.Add(matrix3)
// 	e := Matrix{
// 		{4, 7, 4},
// 		{6, 8, 8},
// 	}

// 	for row := range e {
// 		for col := range e[row] {
// 			val, err := m.Value(row, col)
// 			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
// 			assert.Equal(t, nil, err, "error should be nil")
// 		}
// 	}
// }

// func TestAddIncompatible(t *testing.T) {
// 	_, err := matrix1.Add(matrix2)
// 	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
// }

// func TestSubtract(t *testing.T) {
// 	m, _ := matrix3.Subtract(matrix1)
// 	e := Matrix{
// 		{2, 3, -2},
// 		{2, 0, -2},
// 	}

// 	for row := range e {
// 		for col := range e[row] {
// 			val, err := m.Value(row, col)
// 			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
// 			assert.Equal(t, nil, err, "error should be nil")
// 		}
// 	}
// }

// func TestSubtractIncompatible(t *testing.T) {
// 	_, err := matrix1.Subtract(matrix2)
// 	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
// }

// func TestMultiply(t *testing.T) {
// 	m, _ := matrix3.Multiply(matrix1)
// 	e := Matrix{
// 		{3, 10, 3},
// 		{8, 16, 15},
// 	}

// 	for row := range e {
// 		for col := range e[row] {
// 			val, err := m.Value(row, col)
// 			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
// 			assert.Equal(t, nil, err, "error should be nil")
// 		}
// 	}
// }

// func TestMultiplyIncompatible(t *testing.T) {
// 	_, err := matrix1.Multiply(matrix2)
// 	assert.Equal(t, errors.New("incompatible matrices"), err, "error should be 'incompatible matrices'")
// }

// func TestScale(t *testing.T) {
// 	m := matrix1.Scale(4)
// 	e := Matrix{
// 		{4, 8, 12},
// 		{8, 16, 20},
// 	}

// 	for row := range e {
// 		for col := range e[row] {
// 			val, err := m.Value(row, col)
// 			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
// 			assert.Equal(t, nil, err, "error should be nil")
// 		}
// 	}
// }

// func TestTranspose(t *testing.T) {
// 	m := matrix1.Transpose()
// 	e := Matrix{
// 		{1, 2},
// 		{2, 4},
// 		{3, 5},
// 	}

// 	for row := range e {
// 		for col := range e[row] {
// 			val, err := m.Value(row, col)
// 			assert.Equal(t, e[row][col], val, fmt.Sprintf("Value at position (%v, %v) should be %v", row, col, e[row][col]))
// 			assert.Equal(t, nil, err, "error should be nil")
// 		}
// 	}
// }

// func TestReshape(t *testing.T) {
// 	m := Matrix{
// 		{1, 3, 2, 5, 7, 1},
// 		{6, 8, 3, 3, 7, 9},
// 		{7, 9, 3, 1, 9, 0},
// 		{3, 3, 7, 0, 1, 7},
// 	}
// 	r, err := m.Reshape(2, 12)
// 	e := Matrix{
// 		{1, 3, 2, 5, 7, 1, 6, 8, 3, 3, 7, 9},
// 		{7, 9, 3, 1, 9, 0, 3, 3, 7, 0, 1, 7},
// 	}

// 	assert.Equal(t, true, r.Equal(e), "matrices do not match")
// 	assert.Equal(t, nil, err, "error should be nil")

// 	r, err = r.Reshape(6, 4)
// 	e = Matrix{
// 		{1, 3, 2, 5},
// 		{7, 1, 6, 8},
// 		{3, 3, 7, 9},
// 		{7, 9, 3, 1},
// 		{9, 0, 3, 3},
// 		{7, 0, 1, 7},
// 	}

// 	assert.Equal(t, true, r.Equal(e), "matrices do not match")
// 	assert.Equal(t, nil, err, "error should be nil")
// }

// func TestReshapeIncompatible(t *testing.T) {
// 	_, err := matrix1.Reshape(2, 2)
// 	assert.Equal(t, errors.New("dimensions do not match"), err,
// 		"error should be 'dimensions do not match'")
// }
