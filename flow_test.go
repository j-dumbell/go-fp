package fp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlow2(t *testing.T) {
	actual := Flow2(addOne, square)(2)
	expected := 9
	assert.Equal(t, expected, actual)
}

func TestFlow3(t *testing.T) {
	actual := Flow3(isHello, toInt, addOne)("hello")
	expected := 2
	assert.Equal(t, expected, actual)
}
