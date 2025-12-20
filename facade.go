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

// 注册常来容器
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
