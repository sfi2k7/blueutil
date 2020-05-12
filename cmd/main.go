package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

type myerror struct {
	Name string
}

func (m *myerror) Error() string {
	return "I had some error " + m.Name
}

func main() {
	e := getError()
	var single myerror
	ok := xerrors.As(e, &single)
	fmt.Println(ok, single)
	fmt.Println(e)
	// s := blueutil.NewMem()
	// s.Add("1", "One")
	// s2 := blueutil.NewMemWithParent(s)
	// s2.Add("2", "Two")
	// defer s.Close()
	// defer s2.Close()
	// fmt.Println(s2.Get("2"))
	// fmt.Println(s2.Get("1"))
}

func getError() error {
	return &myerror{Name: "Test"}
	// e := xerrors.New("some text")
	// return e
}
