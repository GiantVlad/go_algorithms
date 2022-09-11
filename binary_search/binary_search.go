package binary_search

func BSearch(needle int, list []int) int {
	low := 0
	high := len(list)
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == needle {
			return mid
		} else if list[mid] <= needle {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
