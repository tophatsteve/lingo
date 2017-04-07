package lingo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddScalarToScalar(t *testing.T) {
	s1 := Scalar(1)
	s2 := Scalar(2)
	s3, err := Add(s1, s2)

	assert.Equal(t, nil, err, "err is not nil")
	assert.Equal(t, Scalar(3), s3, "value of 1 + 2 is not 3")
}
