package inserting_sort

func InsSort(list []int) {
	k := 0
	for i, _ := range list {
		k = i - 1
		for k > -1 {
			if list[k+1] < list[k] {
				list[k], list[k+1] = list[k+1], list[k]
			}
			k--
		}
	}
}
