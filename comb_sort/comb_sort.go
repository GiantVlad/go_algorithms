package comb_sort

func CombSort(list []int) {
	d := int(float64(len(list)) / 1.247)
	var caps []int
	for d > 0 {
		caps = append(caps, d)
		d = int(float64(d) / 1.247)
	}
	for _, cp := range caps {
		k := 0
		for i := 0; i < len(list); i += cp {
			k = i - cp
			for k > -1 {
				if list[k+cp] < list[k] {
					list[k], list[k+cp] = list[k+cp], list[k]
				}
				k -= cp
			}
		}
	}
}
