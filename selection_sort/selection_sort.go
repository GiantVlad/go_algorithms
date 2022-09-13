package selection_sort

func SelSort(list []int) {
	if len(list) < 2 {
		return
	}
	min := 0
	for i, _ := range list {
		min = i
		for j := i + 1; j < len(list); j++ {
			if list[min] > list[j] {
				min = j
			}
		}
		list[i], list[min] = list[min], list[i]
	}
}
