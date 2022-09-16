package shell_sort

func ShellSort(list []int) {
	d := len(list) / 2
	var caps []int
	for d > 0 {
		caps = append(caps, d)
		d = d / 2
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
