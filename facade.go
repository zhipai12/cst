package cst

import (
	"sync"

	"github.com/rrzu/cst/common"
)

type (
	// Typ 类型名称
	Typ[T IConstVal, S any] string

	// FilterFns 过滤函数
	FilterFns[T IConstVal, S any] func(word Word[T, S]) bool
)

// 注册常量容器
var containerSlice sync.Map

// Register 注册常量
func Register[T IConstVal, S any](typ Typ[T, S], cm *Cst[T, S]) *Cst[T, S] {
	containerSlice.Store(typ, cm)

	return cm
}

// Get 获取常量
func Get[T IConstVal, S any](typ Typ[T, S]) *Cst[T, S] {
	v, ok := containerSlice.Load(typ)
	if !ok {
		return nil
	}

	result, ok := v.(*Cst[T, S])
	if !ok {
		return nil
	}

	return result
}

// GetFilterGroup 获取常量并过滤分组
func GetFilterGroup[T IConstVal, S any](typ Typ[T, S], groupNames ...GroupName) *Cst[T, S] {
	v, ok := containerSlice.Load(typ)
	if !ok {
		return nil
	}

	var result = &Cst[T, S]{}
	if len(groupNames) > 0 {
		for _, w := range v.(*Cst[T, S]).Words {
			for _, groupName := range groupNames {
				if w.Group != nil && !common.InSlice(*w.Group, groupName) {
					result.Words = append(result.Words, w)
				}
			}
		}
	}

	if len(result.Words) > 0 {
		result.Typ = v.(*Cst[T, S]).Typ
	}

	return result
}

// GetFilters 获取常量并根据filterFns过滤
func GetFilters[T IConstVal, S any](typ Typ[T, S], filterFns ...FilterFns[T, S]) *Cst[T, S] {
	v, ok := containerSlice.Load(typ)
	if !ok {
		return nil
	}

	result, ok := v.(*Cst[T, S])
	if !ok {
		return nil
	}

	if len(filterFns) > 0 {
		for i, w := range v.(*Cst[T, S]).Words {
			for _, through := range filterFns {
				if !through(w) {
					result.Words = append(result.Words[:i], result.Words[i+1:]...)
				}
			}
		}
	}

	return result
}

// CstWithGroup 保留包含 groupName 的常量
func CstWithGroup[T IConstVal, S any](raw *Cst[T, S], groupNames ...GroupName) (new *Cst[T, S]) {
	if raw == nil {
		return &Cst[T, S]{}
	}

	new = &Cst[T, S]{}

	// 如果没有指定分组，则返回空集合（或根据业务需求可返回原始集合）
	if len(groupNames) == 0 {
		return new
	}

	// 创建分组名称映射，提高查找效率
	groupMap := make(map[GroupName]bool, len(groupNames))
	for _, name := range groupNames {
		groupMap[name] = true
	}

	// 遍历所有常量
	for _, w := range raw.Words {
		// 保留分组在指定列表中的常量
		if w.Group != nil {
			// 检查字段的分组列表中是否包含任意一个指定的分组
			for _, fieldGroup := range *w.Group {
				if groupMap[fieldGroup] {
					new.Words = append(new.Words, w)
					break // 找到一个匹配就跳出，避免重复添加
				}
			}
		}
	}

	// 如果新集合不为空，复制类型信息
	if len(new.Words) > 0 {
		new.Typ = raw.Typ
	}

	return
}
