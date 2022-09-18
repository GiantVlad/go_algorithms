package merge_sort

func MergeSort(list []int) []int {
	if len(list) == 1 {
		return list
	}
	n := len(list) / 2
	right := list[:n]
	left := list[n:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(listL []int, listR []int) (res []int) {
	for len(listL) > 0 && len(listR) > 0 {
		if listL[0] < listR[0] {
			res = append(res, listL[0])
			listL = listL[1:]
		} else {
			res = append(res, listR[0])
			listR = listR[1:]
		}
	}
	if len(listL) > 0 {
		res = append(res, listL...)
	}
	if len(listR) > 0 {
		res = append(res, listR...)
	}
	return
}
