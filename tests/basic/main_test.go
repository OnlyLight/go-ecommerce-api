package basic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), output = %d, actual = %d", input, output, actual)
	// }

	assert.Equal(t, 3, AddOne(2), "AddOne(2) should be 3")
}
