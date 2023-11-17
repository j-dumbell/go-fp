package fp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapI(t *testing.T) {
	fn := func(i int, j int) int {
		return i + j
	}
	actual := MapI(fn)([]int{2, 0, 1})
	expected := []int{2, 1, 3}
	assert.Equal(t, actual, expected)
}

func TestMap(t *testing.T) {
	actual := Map(square)([]int{2, 0, 1})
	expected := []int{4, 0, 1}
	assert.Equal(t, actual, expected)
}

func TestFilterI(t *testing.T) {
	fn := func(i int) bool {
		return i > 0
	}
	actual := Filter(fn)([]int{2, -5, 10})
	expected := []int{2, 10}
	assert.Equal(t, actual, expected)
}

func TestFilter(t *testing.T) {
	actual := Filter(isHello)([]string{"blah", "hello"})
	expected := []string{"hello"}
	assert.Equal(t, actual, expected)
}
