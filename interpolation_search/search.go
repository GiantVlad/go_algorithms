package interpolation_search

func InterpolSearch(needle int, list []int) int {
	low, high := 0, len(list)
	min, max := list[0], list[high-1]

	var guess int

	for {
		if needle < min || needle > max {
			return -1
		}

		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(needle-min) / float64(max-min)))
			guess = low + offset
		}

		if list[guess] == needle {
			for guess > 0 && list[guess-1] == needle {
				guess--
			}
			return guess
		} else if list[guess] > needle {
			high = guess - 1
			max = list[high]
		} else {
			low = guess + 1
			min = list[low]
		}
	}
}
