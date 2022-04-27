package main

import (
	"fmt"

	"github.com/jasonfriedland/generics/pkg/set"
)

func main() {
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
}
