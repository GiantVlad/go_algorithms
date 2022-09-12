package buble_sort

func Sort(list []int) {
	if len(list) < 2 {
		return
	}
	for k := len(list) - 1; k > 0; k-- {
		for n := 0; n < k; n++ {
			if list[n] > list[n+1] {
				list[n], list[n+1] = list[n+1], list[n]
			}
		}
	}
	return
}
