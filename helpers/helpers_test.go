package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPrimes(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	result := FindPrimes(30)

	assert.Equal(t, expected, result)
}
