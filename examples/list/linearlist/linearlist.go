package main

import (
	"fmt"

	"github.com/icheontao/dsl/list/linearlist"
)

func main() {

	l1 := linearlist.New()
	l1.Append("a", "b", "c", "d", "b")
	fmt.Println(l1) // LinearList([a, b, c, d, b])
	a := l1.Len()
	fmt.Println(a) // 5
	b, ok := l1.Get(2)
	fmt.Println(b, ok) // c false
	c, ok := l1.Get(6)
	fmt.Println(c, ok) // <nil> false
	d := l1.Indexof("b")
	fmt.Println(d) // 1
	l1.Insert(2, "f", "g", "h")
	fmt.Println(l1) // LinearList([a, b, f, c, d, b])
	l1.Remove("b")
	fmt.Println(l1) // LinearList([a, f, c, d, b])
	e := l1.Empty()
	fmt.Println(e) // false
	l1.Clear()
	f := l1.Len()
	fmt.Println(f) // 0
	g := l1.Empty()
	fmt.Println(g) // true

}
