package stack

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackUsage(t *testing.T) {
	s := New(4)
	assert.Equal(t, 0, s.Len())

	assert.Equal(t, true, s.Push("salad"))
	assert.Equal(t, 1, s.Len())

	assert.Equal(t, true, s.Push("pizza"))
	assert.Equal(t, true, s.Push("wings"))
	assert.Equal(t, true, s.Push("brownie"))
	assert.Equal(t, false, s.Push("soda"))

	a, ok := s.Pop()
	assert.Equal(t, "brownie", a)
	assert.True(t, ok)

	a, ok = s.Pop()
	assert.Equal(t, "wings", a)
	assert.True(t, ok)

	a, ok = s.Pop()
	assert.Equal(t, "pizza", a)
	assert.True(t, ok)
	assert.Equal(t, 1, s.Len())

	assert.Equal(t, true, s.Push("soda"))
}

func TestPeekPopEmpty(t *testing.T) {
	s := New(4)
	a, ok := s.Peek()
	assert.Equal(t, nil, a)
	assert.False(t, ok)
	a, ok = s.Pop()
	assert.Equal(t, nil, a)
	assert.False(t, ok)
}

func TestPeekPopNotEmpty(t *testing.T) {
	s := New(4)
	s.Push(randomInt())
	ak, okk := s.Peek()
	assert.True(t, okk)
	ap, okp := s.Pop()
	assert.Equal(t, ak, ap)
	assert.Equal(t, okk, okp)
}

// Exercise a random walk of operations, and match that to an expected net result.
func TestZeroKnowledgeAccess(t *testing.T) {
	s := New(randomInt() + 20)
	runningTotal := 0
	for remaining := randomInt() + 100; remaining > 0; remaining-- {
		if randomShouldPush() {
			num := randomInt()
			expectSuccess := s.Len() < s.Cap()
			if expectSuccess {
				runningTotal += num
			}
			actualSuccess := s.Push(num)
			assert.Equal(t, expectSuccess, actualSuccess)
			if actualSuccess {
				top, _ := s.Peek()
				assert.Equal(t, num, top)
			}
		} else {
			expectSuccess := s.Len() > 0
			peekNum, peekSuccess := s.Peek()
			popNum, popSuccess := s.Pop()
			assert.Equal(t, expectSuccess, peekSuccess)
			assert.Equal(t, expectSuccess, popSuccess)
			assert.Equal(t, peekNum, popNum)
			if popSuccess {
				runningTotal -= popNum.(int)
			}
		}
	}
	eventualTotal := 0
	for eventualNum, ok := s.Pop(); ok; eventualNum, ok = s.Pop() {
		eventualTotal += eventualNum.(int)
	}
	assert.Equal(t, runningTotal, eventualTotal)
}

func randomInt() int {
	return int(rand.Int31n(100))
}

func randomShouldPush() bool {
	return rand.Int31n(2) == 1
}
