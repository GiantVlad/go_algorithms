package linear_search

func Search(datalist []int, targ int) int {
	for idx, val := range datalist {
		if val == targ {
			return idx
		}
	}
	return -1
}
