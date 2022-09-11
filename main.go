package main

import (
	"fmt"
	"linearSearch/linearSearch"
)

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	idx := linearSearch.Search(items, 58)
	if idx > -1 {
		fmt.Printf("Found element, idx: %d", idx)
		return
	}
	fmt.Println("Not found")
}