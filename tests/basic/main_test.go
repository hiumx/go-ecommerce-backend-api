package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	var (
		input    = 1
		expected = 2
	)

	// actual := AddOne(input)
	// if actual != expected {
	// 	t.Errorf("AddOne(%d) = %d, Actual: %d, Expected: %d", input, actual, actual, expected)
	// }

	assert.Equal(t, expected, AddOne(input), "Not pass")
	assert.NotEqual(t, 3, 4, "")
}

func TestAddOne2(t *testing.T) {
	var (
		input    = 1
		expected = 2
	)

	// actual := AddOne(input)
	// if actual != expected {
	// 	t.Errorf("AddOne(%d) = %d, Actual: %d, Expected: %d", input, actual, actual, expected)
	// }

	assert.Equal(t, expected, AddOne2(input), "Not pass")
	assert.NotEqual(t, 3, 4, "")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not exec")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Exec")
}
