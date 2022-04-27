package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	// String test
	t.Run("Set, string", func(t *testing.T) {
		assert := assert.New(t)
		s := New[string]()

		s.Add("hello")
		s.Add("hello")
		assert.Equal(s.Len(), 1)
		assert.Equal(s.In("hello"), true)

		s.Add("foo")
		s.Add("bar")
		assert.Equal(s.Len(), 3)
		assert.Equal(s.In("hello") && s.In("foo") && s.In("bar"), true)

		s.Del("hello")
		assert.Equal(s.Len(), 2)
		assert.Equal(s.In("hello"), false)
	})
	// Int test
	t.Run("Set, int", func(t *testing.T) {
		assert := assert.New(t)
		s := New[int]()

		s.Add(123)
		s.Add(123)
		assert.Equal(s.Len(), 1)
		assert.Equal(s.In(123), true)

		s.Add(456)
		s.Add(789)
		assert.Equal(s.Len(), 3)
		assert.Equal(s.In(123) && s.In(456) && s.In(789), true)

		s.Del(123)
		assert.Equal(s.Len(), 2)
		assert.Equal(s.In(123), false)
	})
	// Struct test
	t.Run("Set, struct", func(t *testing.T) {
		assert := assert.New(t)
		type tmp struct {
			s string
			f float64
			b bool
		}
		s := New[tmp]()

		s.Add(tmp{"hello", 12.0, true})
		s.Add(tmp{"hello", 12.0, true})
		assert.Equal(s.Len(), 1)
		assert.Equal(s.In(tmp{"hello", 12.0, true}), true)

		s.Add(tmp{"goodbye", 3.14159, false})
		s.Add(tmp{"hello", 12.0, false})
		assert.Equal(s.In(tmp{"hello", 12.0, true}) &&
			s.In(tmp{"goodbye", 3.14159, false}) &&
			s.In(tmp{"hello", 12.0, false}), true)

		s.Del(tmp{"hello", 12.0, true})
		assert.Equal(s.In(tmp{"hello", 12.0, true}), false)
	})
}

func benchmarkAddInt(b *testing.B, size int) {
	s := New[int]()
	for i := 0; i < size; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Add(-1)
	}
}

func BenchmarkAddInt1(b *testing.B)       { benchmarkAddInt(b, 1) }
func BenchmarkAddInt10(b *testing.B)      { benchmarkAddInt(b, 10) }
func BenchmarkAddInt1000(b *testing.B)    { benchmarkAddInt(b, 1000) }
func BenchmarkAddInt1000000(b *testing.B) { benchmarkAddInt(b, 1000000) }

func benchmarkDelInt(b *testing.B, size int) {
	s := New[int]()
	for i := 0; i < size; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Del(1)
	}
}

func BenchmarkDelInt1(b *testing.B)       { benchmarkDelInt(b, 1) }
func BenchmarkDelInt10(b *testing.B)      { benchmarkDelInt(b, 10) }
func BenchmarkDelInt1000(b *testing.B)    { benchmarkDelInt(b, 1000) }
func BenchmarkDelInt1000000(b *testing.B) { benchmarkDelInt(b, 1000000) }
