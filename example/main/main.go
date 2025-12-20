package main

import (
	"encoding/json"
	"fmt"

	"github.com/rrzu/cst"
	"github.com/rrzu/cst/example"
)

func main() {
	var a []byte

	// a, _ = json.Marshal(cst.Get(example.TypCPActivityStatus).ToOptions())
	// fmt.Println(string(a))
	// a, _ = json.Marshal(cst.Get(example.TypBasicIndicator).ToOptions())
	// fmt.Println(string(a))
	// a, _ = json.Marshal(cst.Get(example.TypShowUnit).ToOptions())
	// fmt.Println(string(a))
	//
	// a, _ = json.Marshal(cst.Get(example.TypDB))
	// fmt.Println(string(a))

	fmt.Println("here")
	a, _ = json.Marshal(cst.Get(example.TypDB))
	fmt.Println("1:", string(a))

	// a, _ = json.Marshal(cst.GetFilterGroup(example.TypDB, "OLAP"))
	// fmt.Println("2:", string(a))
	//
	// a, _ = json.Marshal(cst.Get(example.TypDB))
	// fmt.Println("3:", string(a))
	//
	// a, _ = json.Marshal(cst.GetFilterGroup(example.TypDB, "OLTP"))
	// fmt.Println("4:", string(a))
}
