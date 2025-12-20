package example

import (
	"fmt"

	"github.com/rrzu/cst"
)

// 活动状态
const (
	CPActivityStatusNone      CPActivityStatus = iota // none
	CPActivityStatusDraft                             // 草稿
	CPActivityStatusWaitStart                         // 待开始
	CPActivityStatusWorking                           // 进行中
	CPActivityStatusOver                              // 已结束
)

// CPActivityStatus 活动状态
type CPActivityStatus uint8

var cPActivityStatus = &cst.Cst[CPActivityStatus, any]{
	Typ: cst.DataTypeNumber,
	Words: cst.Words[CPActivityStatus, any]{
		{
			Value:  CPActivityStatusNone,
			CnName: "none",
		},
		{
			Value:  CPActivityStatusDraft,
			CnName: "草稿",
		},
		{
			Value:  CPActivityStatusWaitStart,
			CnName: "待开始",
		},
		{
			Value:  CPActivityStatusWorking,
			CnName: "进行中",
		},
		{
			Value:  CPActivityStatusOver,
			CnName: "已结束",
			Group:  &cst.Group{"homePage"},
		},
	},
}

func init() {
	fmt.Println("init cPActivityStatus")
	cst.Register(TypCPActivityStatus, cPActivityStatus)
}
