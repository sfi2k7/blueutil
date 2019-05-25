package main

import (
	"fmt"

	"github.com/sfi2k7/blueutil"
)

func main() {
	s := blueutil.NewMem()
	s.Add("1", "One")
	s2 := blueutil.NewMemWithParent(s)
	s2.Add("2", "Two")
	defer s.Close()
	defer s2.Close()
	fmt.Println(s2.Get("2"))
	fmt.Println(s2.Get("1"))
}
