package fp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurry2(t *testing.T) {
	fn := func(i int, s string) string {
		return fmt.Sprintf("%s %v", s, i)
	}
	actual := Curry2(fn)(2)("blah")
	expected := "blah 2"

	assert.Equal(t, expected, actual)
}

func TestCurry3(t *testing.T) {
	fn := func(i, j, k int) int {
		return (i * i) + j - k
	}
	actual := Curry3(fn)(3)(1)(5)
	expected := 5

	assert.Equal(t, expected, actual)
}
