package main

import (
	"fmt"
	"linearSearch/hanoi_tower"
)

func main() {
	beadNum := 5 // This is the initial number of beads
	fmt.Printf("This is a Hannukah game with %d beads \n\r", beadNum)
	count := hanoi_tower.Hanoi(beadNum, "A", "B", "C")
	fmt.Printf("Game over: %d steps spent in total \n\r", count)
}
