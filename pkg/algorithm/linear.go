package algorithm

type TransformingCallback[TSource, TDest any] func(el TSource) TDest

func Transformed[TSource, TDest any](source []TSource, f TransformingCallback[TSource, TDest]) []TDest {
	dest := make([]TDest, len(source))
	for i, el := range source {
		dest[i] = f(el)
	}
	return dest
}
