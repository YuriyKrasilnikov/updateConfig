package utils

// универсальный обход дерева
func traversalTreeBF[T any](arr []*T, app func(*T) bool, getNext func(*T) []*T) {
	for len(arr) > 0 {
		nextArr := []*T{}
		for _, elm := range arr {
			app(elm)
			nextArr = append(nextArr, getNext(elm)...)
		}
		arr = nextArr
	}
}
