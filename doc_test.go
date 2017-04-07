package lingo

import "fmt"

// Add Scalar s1 to Scalar s2, creating Scalar s3
func ExampleScalar_Add() {
	s1 := Scalar(1)
	s2 := Scalar(2)
	s3, _ := Add(s1, s2)

	fmt.Println(s3)
	// Output: 3
}
