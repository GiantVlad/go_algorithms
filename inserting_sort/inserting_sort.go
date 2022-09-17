package inserting_sort

func InsSort(list []int) {
	for i := 1; i < len(list); i++ {
		j := i - 1
	inserted:
		for j > -1 {
			if list[i] > list[j] {
				insert(list, i, j+1)
				break inserted
			}
			if j == 0 {
				insert(list, i, j)
				break inserted
			}
			j--
		}
	}
}

func insert(list []int, from int, to int) {
	for i := from; i > to; i-- {
		list[i-1], list[i] = list[i], list[i-1]
	}
}
