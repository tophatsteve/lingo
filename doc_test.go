package lingo

// import "fmt"

// // Add matrix m1 and matrix m2, creating matrix m3.
// func ExampleMatrix_Add() {
// 	m1 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}

// 	m2 := Matrix{
// 		{3, 5, 1},
// 		{4, 4, 3},
// 	}

// 	m3, _ := m1.Add(m2)

// 	fmt.Println(m3)
// 	// Output: [4 7 4]
// 	// [6 8 8]
// }

// // Subtract matrix m1 from matrix m2, creating matrix m3.
// func ExampleMatrix_Subtract() {
// 	m1 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}

// 	m2 := Matrix{
// 		{3, 5, 1},
// 		{4, 4, 3},
// 	}

// 	m3, _ := m2.Subtract(m1)

// 	fmt.Println(m3)
// 	// Output: [2 3 -2]
// 	// [2 0 -2]
// }

// // Element wise multiplcation of matrix m1 by matrix m2, creating matrix m3.
// func ExampleMatrix_Multiply() {
// 	m1 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}

// 	m2 := Matrix{
// 		{3, 5, 1},
// 		{4, 4, 3},
// 	}

// 	m3, _ := m1.Multiply(m2)

// 	fmt.Println(m3)
// 	// Output: [3 10 3]
// 	// [8 16 15]
// }

// // Dot product of matrix m1 and matrix m2, creating matrix m3.
// func ExampleMatrix_Dot() {
// 	m1 := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}

// 	m2 := Matrix{
// 		{1, 2},
// 		{3, 2},
// 		{4, 5},
// 	}

// 	m3, _ := m1.Dot(m2)

// 	fmt.Println(m3)
// 	// Output: [19 21]
// 	// [34 37]
// }

// // Get the value at position 1,2 in matrix m.
// func ExampleMatrix_Value() {
// 	m := Matrix{
// 		{1, 2, 3},
// 		{2, 4, 5},
// 	}

// 	v, _ := m.Value(1, 2)

// 	fmt.Println(v)
// 	// Output: 5
// }

// // Rows
// // Columns
// // Dimensions
// // Equal
// // Scale
// // Transpose
// // Reshape
// // ToVector
