package example

import (
	"fmt"

	"github.com/rrzu/cst"
)

// 展示单位
const (
	DBNone DB = iota
	DBMysql
	DBPostgres
	DBOracle
	DBHologres
	DBClickhouse
)

// DB 展示单位
type DB int

type MineDB *cst.Cst[Database, any]

var db = &cst.Cst[DB, cst.Options[Database, any]]{
	Typ: cst.DataTypeNumber,
	Words: cst.Words[DB, cst.Options[Database, any]]{
		{
			Value:  DBNone,
			CnName: "无",
		},
		{
			Value:  DBMysql,
			Group:  &cst.Group{"OLTP"},
			CnName: "MySQL",
			Mine:   database.ToOptions(),
		},
		{
			Value:  DBPostgres,
			Group:  &cst.Group{"OLTP"},
			CnName: "Postgres",
		},
		{
			Value:  DBOracle,
			Group:  &cst.Group{"OLTP"},
			CnName: "Oracle",
		},
		{
			Value:  DBHologres,
			Group:  &cst.Group{"OLAP"},
			CnName: "Hologres",
		},
		{
			Value:  DBClickhouse,
			Group:  &cst.Group{"OLAP", "TEST"},
			CnName: "Clickhouse",
		},
	},
}

func init() {
	cst.Register(TypDB, db)

	newDb := cst.CstWithGroup(db, "TEST")
	fmt.Printf("newDb: %+v\n", newDb)
}
