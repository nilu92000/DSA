func processSingle(w http.ResponseWriter, r *http.Request) {
	process(w, r, false)
}

func sequentialSort(toSort [][]int) [][]int {
	var sortedLists [][]int

	for _, innerList := range toSort {
		sortedInnerList := make([]int, len(innerList))
		copy(sortedInnerList, innerList)
		sort.Ints(sortedInnerList)
		sortedLists = append(sortedLists, sortedInnerList)
	}

	return sortedLists
}
