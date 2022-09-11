package main

import (
	"fmt"
	"linearSearch/linear_search"
)

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	idx := linear_search.Search(items, 58)
	if idx > -1 {
		fmt.Printf("Found element, idx: %d", idx)
		return
	}
	fmt.Println("Not found")
}
