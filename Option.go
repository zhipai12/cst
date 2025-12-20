package cst

// 选项数据结构
type (
	Options[T IConstVal, S any] struct {
		Typ  DataType       `json:"typ"`
		Opts []Option[T, S] `json:"opts"`
	}

	Option[T IConstVal, S any] struct {
		CnName  string         `json:"cn_name"`
		Val     interface{}    `json:"val"`
		Mine    S              `json:"mine"`
		Sub     *Options[T, S] `json:"sub,omitempty"`
		Cascade interface{}    `json:"cascade,omitempty"`
	}
)
