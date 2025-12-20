package example

import "github.com/rrzu/cst"

const (
	TypCPActivityStatus cst.Typ[CPActivityStatus, any]              = "CPActivityStatus"
	TypBasicIndicator   cst.Typ[BasicIndicator, MineBasicIndicator] = "BasicIndicator"
	TypShowUnit         cst.Typ[ShowUnit, any]                      = "ShowUnit"
	TypDB               cst.Typ[DB, cst.Options[Database, any]]     = "DB"
	TypDatabase         cst.Typ[Database, any]                      = "Database"
)
