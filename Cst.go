package cst

import (
	"fmt"
	"sync"

	"github.com/rrzu/cst/common"
)

// 变量内存缓存（读多写少 key:*Cst[T, S] value:map[T]Word[T, S]）
var containerMap sync.Map

type (
	// Cst 定义
	Cst[T IConstVal, S any] struct {
		Typ   DataType    // 数据类型
		Words Words[T, S] // 字段集合
	}

	// Words 字段集合
	Words[T IConstVal, S any] []Word[T, S]

	// Word 字段
	Word[T IConstVal, S any] struct {
		Value  T             // 值
		Group  *Group        // 使用分组
		CnName string        // 中文名
		Sub    *[]Word[T, S] // 子字段
		Mine   S             // 附加信息
	}

	Group     []GroupName // 使用分区
	GroupName string
)

// ToOptions 转换为选项数据格式
func (c *Cst[T, S]) ToOptions(groupNames ...GroupName) Options[T, S] {
	opts := make([]Option[T, S], 0)
	for _, mp := range c.Words {
		if len(groupNames) > 0 {
			if mp.Group == nil {
				continue
			}

			found := false
			for _, groupName := range groupNames {
				if !common.InSlice(*mp.Group, groupName) {
					continue
				}
				found = true
			}

			if !found {
				continue
			}
		}

		var option = Option[T, S]{
			CnName: mp.CnName,
			Val:    mp.Value,
			Mine:   mp.Mine,
		}

		if mp.Sub != nil {
			x := Cst[T, S]{Typ: c.Typ, Words: *mp.Sub}

			sub := x.ToOptions()
			option.Sub = &sub
		}

		opts = append(opts, option)
	}

	return Options[T, S]{
		Typ:  c.Typ,
		Opts: opts,
	}
}

// ToWordMap 将 Cst.Words 转换为 map[T]Word[T].
func (c *Cst[T, S]) ToWordMap() (Cst map[T]Word[T, S]) {
	// 根据地址写入内存
	ptrKey := fmt.Sprintf("%p", c)
	if v, ok := containerMap.Load(ptrKey); ok {
		return v.(map[T]Word[T, S])
	} else {
		defer func() {
			containerMap.Store(ptrKey, Cst)
		}()
	}

	Cst = make(map[T]Word[T, S])
	for _, item := range c.Words {
		Cst[item.Value] = item
	}
	return Cst
}

// IsAllValid 验证值是否在 Cst.Words 中
func (c *Cst[T, S]) IsAllValid(value []T) bool {
	mp := c.ToWordMap()
	if mp == nil || value == nil || len(value) == 0 {
		return false
	}

	for _, t := range value {
		if _, ok := mp[t]; !ok {
			return false
		}
	}

	return true
}

// IsValid 验证值是否在 Cst.Words 中
func (c *Cst[T, S]) IsValid(value T) bool {
	mp := c.ToWordMap()
	if mp == nil {
		return false
	}

	if _, ok := mp[value]; !ok {
		return false
	}

	return true
}
