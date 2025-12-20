package example

import (
	"fmt"

	"github.com/rrzu/cst"
)

const (
	DatabaseZulin   Database = "zulin"
	DatabaseServer  Database = "server"
	DatabaseProduct Database = "product"
)

type Database string

var database = &cst.Cst[Database, any]{
	Typ: cst.DataTypeString,
	Words: cst.Words[Database, any]{
		{
			Value:  DatabaseZulin,
			CnName: "zulin",
		},
		{
			Value:  DatabaseServer,
			CnName: "server",
		},
		{
			Value:  DatabaseProduct,
			CnName: "product",
		},
	},
}

func init() {
	fmt.Println("init database")
	cst.Register(TypDatabase, database)
}
