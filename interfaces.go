package cst

type (
	// IConstVal 常量取值范围
	IConstVal interface {
		~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64 |
			~int8 | ~int16 | ~int32 | ~int | ~int64 |
			~float32 | ~float64 |
			~string
	}

	ICst[T IConstVal, S any] interface {
		ToOptions() Options[T, S]          // 转换为选项数据格式
		ToWordMap() (ret map[T]Word[T, S]) // 将 Cst.Words 转换为 map[T]Word[T].
		IsAllValid(value []T) bool         // 验证值是否在 Cst.Words 中
		IsValid(value T) bool              // 验证值是否在 Cst.Words 中
	}
)
