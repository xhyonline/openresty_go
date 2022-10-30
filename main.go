package main

import (
	"C"
	"fmt"
	"github.com/xhyonline/openresty_go/lib"
)

func main() {
	lib.Add("/", 1)
	lib.Add("/a.css", 2)
	lib.Add("/d/a.css", 3)

	fmt.Println(lib.Walk("/d/a.css"))
}

//export Add
func Add() {
	
}
