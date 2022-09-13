package quick_sort

import "math/rand"

func QSort(list []int) {
	if len(list) < 2 {
		return
	}
	pivot := rand.Int() % len(list)

	left, right := 0, len(list)-1
	list[right], list[pivot] = list[pivot], list[right]

	for i, _ := range list {
		if list[i] < list[right] {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}
	list[right], list[left] = list[left], list[right]
	QSort(list[:left])
	QSort(list[left+1:])
}
