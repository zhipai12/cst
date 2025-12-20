package example

import (
	"fmt"

	"github.com/rrzu/cst"

	"github.com/rrzu/cst/common"
)

// 基础指标
const (
	BasicIndicatorDeliveryRate              BasicIndicator = "delivery_rate"                // 发货率
	BasicIndicatorAvgPrepaymentPeriodNum    BasicIndicator = "avg_prepayment_period_num"    // 平均预收期数
	BasicIndicatorAvgValidResponseSecond    BasicIndicator = "avg_valid_response_second"    // 响应平均时效（工作时间）
	BasicIndicatorAvgSentSecond             BasicIndicator = "avg_sent_second"              // 发货平均时效
	BasicIndicatorWaitingDeliveryRate       BasicIndicator = "waiting_delivery_rate"        // 待发货比例
	BasicIndicatorPaymentAttemptSuccessRate BasicIndicator = "payment_attempt_success_rate" // 扣款成功比例
	BasicIndicatorDeliveryWarehouseRate     BasicIndicator = "delivery_warehouse_rate"      // 标记发货比例
	BasicIndicatorAssessmentDeliveryRate    BasicIndicator = "assessment_delivery_rate"     // 考核发货率
)

// BasicIndicator 基础指标
type BasicIndicator string

// MineBasicIndicator 基础指标 Mine
type MineBasicIndicator struct {
	ShowValue func(val float64) float64 `json:"-"`       // 格式化值
	Unit      ShowUnit                  `json:"unit"`    // 格式化单位
	IsHide    bool                      `json:"is_hide"` // 是否隐藏
	Alias     *string                   `json:"alias"`   // 别名
}

var basicIndicator = &cst.Cst[BasicIndicator, MineBasicIndicator]{
	Typ: "",
	Words: cst.Words[BasicIndicator, MineBasicIndicator]{
		{
			Value:  BasicIndicatorDeliveryRate,
			CnName: "发货率",
			Mine: MineBasicIndicator{
				ShowValue: func(val float64) float64 { return common.Round2F(val * 100) },
				Unit:      ShowUnitPercent,
				IsHide:    true,
			},
		},
		{
			Value:  BasicIndicatorAvgPrepaymentPeriodNum,
			CnName: "平均预收期数",
			Mine: MineBasicIndicator{
				ShowValue: func(val float64) float64 { return common.Round2F(val) },
				Unit:      ShowUnitNone,
			},
		},
		{
			Value:  BasicIndicatorAvgValidResponseSecond,
			CnName: "响应平均时效（工作时间）",
			Mine: MineBasicIndicator{
				ShowValue: func(val float64) float64 { return common.Round2F(val / 3600.0) },
				Unit:      ShowUnitTimeH,
			},
		},
		{
			Value:  BasicIndicatorAvgSentSecond,
			CnName: "发货平均时效",
			Mine: MineBasicIndicator{
				ShowValue: func(val float64) float64 { return common.Round2F(val / 3600.0) },
				Unit:      ShowUnitTimeH,
			},
		},
		{
			Value:  BasicIndicatorWaitingDeliveryRate,
			CnName: "待发货比例",
			Mine: MineBasicIndicator{
				ShowValue: func(val float64) float64 { return common.Round2F(val * 100) },
				Unit:      ShowUnitPercent,
			},
		},
		{
			Value:  BasicIndicatorPaymentAttemptSuccessRate,
			CnName: "扣款成功比例",
			Mine: MineBasicIndicator{
				ShowValue: func(val float64) float64 { return common.Round2F(val * 100) },
				Unit:      ShowUnitPercent,
			},
		},
		{
			Value:  BasicIndicatorDeliveryWarehouseRate,
			CnName: "标记发货比例",
			Mine: MineBasicIndicator{
				ShowValue: func(val float64) float64 { return common.Round2F(val * 100) },
				Unit:      ShowUnitPercent,
			},
		},
		{
			Value:  BasicIndicatorAssessmentDeliveryRate,
			CnName: "考核发货率",
			Mine: MineBasicIndicator{
				ShowValue: func(val float64) float64 { return common.Round2F(val * 100) },
				Unit:      ShowUnitPercent,
				Alias:     common.ToPtr("发货率"),
			},
		},
	},
}

func init() {
	fmt.Println("init basicIndicator")
	cst.Register(TypBasicIndicator, basicIndicator)
}
