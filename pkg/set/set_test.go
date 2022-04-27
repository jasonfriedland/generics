package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	// String test
	t.Run("Set, string", func(t *testing.T) {
		s := New[string]()

		s.Add("hello")
		s.Add("hello")
		assert.Equal(t, s.Len(), 1)
		assert.Equal(t, s.items, []string{"hello"})

		s.Add("foo")
		s.Add("bar")
		assert.Equal(t, s.items, []string{"hello", "foo", "bar"})

		s.Del("hello")
		assert.Equal(t, s.items, []string{"bar", "foo"})
	})
	// Int test
	t.Run("Set, int", func(t *testing.T) {
		s := New[int]()

		s.Add(1234)
		s.Add(1234)
		assert.Equal(t, s.Len(), 1)
		assert.Equal(t, s.items, []int{1234})

		s.Add(5678)
		s.Add(9101)
		assert.Equal(t, s.items, []int{1234, 5678, 9101})

		s.Del(1234)
		assert.Equal(t, s.items, []int{9101, 5678})
	})
	// Struct test
	t.Run("Set, struct", func(t *testing.T) {
		type tmp struct {
			s string
			f float64
			b bool
		}
		s := New[tmp]()

		s.Add(tmp{"hello", 12.0, true})
		s.Add(tmp{"hello", 12.0, true})
		assert.Equal(t, s.Len(), 1)
		assert.Equal(t, s.items, []tmp{
			{"hello", 12.0, true},
		})

		s.Add(tmp{"goodbye", 3.14159, false})
		s.Add(tmp{"hello", 12.0, false})
		assert.Equal(t, s.items, []tmp{
			{"hello", 12.0, true},
			{"goodbye", 3.14159, false},
			{"hello", 12.0, false},
		})

		s.Del(tmp{"hello", 12.0, true})
		assert.Equal(t, s.items, []tmp{
			{"hello", 12.0, false},
			{"goodbye", 3.14159, false},
		})
	})
}
