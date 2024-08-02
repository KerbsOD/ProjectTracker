package internal

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func MaxElement[T any](arr []T, compare func(a, b T) bool) T {
	if len(arr) == 0 {
		panic("Array must contain at least one element")
	}
	currentMax := arr[0]
	for _, elem := range arr {
		if compare(elem, currentMax) {
			currentMax = elem
		}
	}
	return currentMax
}

func Max[T any](anElement T, anotherElement T, compare func(a, b T) bool) T {
	if compare(anElement, anotherElement) {
		return anElement
	}
	return anotherElement
}
