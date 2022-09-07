package main

import "fmt"

type A struct {
	A  string `tjson:"名称"`
	ID int64  `tjson:"序号"`
}

func main() {
	a := A{A: "a", ID: 1}
	aa, err := Marshal(a)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(aa))
}
