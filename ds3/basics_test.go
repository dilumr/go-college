package ds3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInterface(t *testing.T) {
	x := "tiger"
	var y interface{} = x
	y = 42
	z := y.(int)
	assert.Equal(t, x, z)
}

func TestFunctionClosure(t *testing.T) {
	state := make([]int, 0, 128)
	appender := func(num int) {
		state = append(state, num)
	}
	callForEach(0, 100, appender)
	assert.Len(t, state, 100)
}

func callForEach(start int, count int, f func(int)) {
	for i := 0; i < count; i++ {
		f(i + start)
	}
}
