package internal

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func MaximizeElementByComparer[T any](arr []T, compare func(a, b T) bool) T {
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

func ArrayWithRepeatedElements[T comparable](arr []T) []T {
	countMap := make(map[T]int)
	for _, element := range arr {
		countMap[element]++
	}

	repeatedElements := []T{}

	for element, count := range countMap {
		if count > 1 {
			repeatedElements = append(repeatedElements, element)
		}
	}

	return repeatedElements
}
