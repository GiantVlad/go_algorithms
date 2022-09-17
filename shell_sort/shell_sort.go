package shell_sort

func ShellSort(list []int) {
	d := len(list) / 2
	var caps []int
	for d > 0 {
		caps = append(caps, d)
		d = d / 2
	}
	for _, cp := range caps {
		for i := cp; i < len(list); i += cp {
			j := i
			for j > 0 {
				if list[j-cp] > list[j] {
					list[j-cp], list[j] = list[j], list[j-cp]
				}
				j = j - cp
			}
		}
	}
}
