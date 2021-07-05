package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_divide(t *testing.T) {

	assert.Equal(t, 1, divide(2, 2))
	assert.Equal(t, 0, divide(2, 4))
	assert.Equal(t, 0, divide(-2, 4))
	assert.Equal(t, -1, divide(-2, 2))
	assert.Equal(t, 2, divide(5, 2))

	assert.Equal(t, 3, divide(10, 3))
	assert.Equal(t, -3, divide(-10, 3))
	assert.Equal(t, 4, divide(12, 3))
	assert.Equal(t, -2, divide(7, -3))
	assert.Equal(t, 1, divide(-1, -1))
	assert.Equal(t, 0, divide(1, 2))
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
	assert.Equal(t, 2147483647, divide(2147483647, 1))
}
