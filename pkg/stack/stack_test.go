package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	// String test
	t.Run("Stack, string", func(t *testing.T) {
		assert := assert.New(t)
		s := New[string]()

		v, ok := s.Pop()
		assert.Equal(v, "")
		assert.Equal(ok, false)

		s.Push("hello")
		s.Push("hello")
		assert.Equal(s.Len(), 2)

		s.Push("foo")
		s.Push("bar")
		assert.Equal(s.Len(), 4)

		v, ok = s.Pop()
		assert.Equal(v, "bar")
		assert.Equal(ok, true)
	})
	// Int test
	t.Run("Stack, int", func(t *testing.T) {
		assert := assert.New(t)
		s := New[int]()

		v, ok := s.Pop()
		assert.Equal(v, 0)
		assert.Equal(ok, false)

		s.Push(42)
		s.Push(73)
		assert.Equal(s.Len(), 2)

		s.Push(1)
		s.Push(880)
		assert.Equal(s.Len(), 4)

		v, ok = s.Pop()
		assert.Equal(v, 880)
		assert.Equal(ok, true)
	})
}
