package main

import (
	"fmt"

	"github.com/jasonfriedland/generics/pkg/set"
)

func main() {
	// Strings
	s1 := set.New[string]()
	s1.Add("hello")
	s1.Add("hello")
	s1.Add("hello")
	s1.Add("foo")
	s1.Add("bar")
	fmt.Println(s1.In("foo"))
	s1.Del("x")
	s1.Del("bar")
	fmt.Println(s1.In("foo"))
	fmt.Println(s1.In("bar"))
	fmt.Println(s1.Len())

	// Ints
	s2 := set.New[int]()
	s2.Add(1)
	s2.Add(43)
	s2.Add(43)
	s2.Add(43)
	s2.Add(43)
	s2.Add(89)
	fmt.Println(s2.In(43))
	s2.Del(10)
	s2.Del(89)
	fmt.Println(s2.In(43))
	fmt.Println(s2.In(89))
	fmt.Println(s2.Len())

	// Structs of comparable types
	type tmp struct {
		s string
		f float64
		b bool
	}
	s3 := set.New[tmp]()
	s3.Add(tmp{"hello", 12.0, true})
	s3.Add(tmp{"hello", 12.0, true})
	s3.Add(tmp{"foo", 3.14, false})
	fmt.Println(s3.In(tmp{"hello", 12.0, true}))
	s3.Del(tmp{"foo", 3.14, false})
	fmt.Println(s3.Len())
}
