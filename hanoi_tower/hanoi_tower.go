package hanoi_tower

import "fmt"

var count = 0

func Hanoi(beadNum int, pillarA string, pillarB string, pillarC string) int {
	if beadNum == 1 {
		// If there is only one bead, move from A to C, game over
		move(beadNum, pillarA, pillarC)

		return count
	}

	// Step 2: move all the plates above N (that is, N-1 judgments) from A to B. At this time, C is the transfer station
	Hanoi(beadNum-1, pillarA, pillarC, pillarB)
	// Step 2: move the Nth plate from A to C
	move(beadNum, pillarA, pillarC)
	// Step 3: move the remaining n-1 disks on B from B to C. At this time, A is the transfer station
	Hanoi(beadNum-1, pillarB, pillarA, pillarC)

	return count
}

func move(beadNum int, pillarSource string, pillarTarget string) {
	count += 1
	fmt.Printf("Step %d: bead of %d from %s move to %s \n\r", count, beadNum, pillarSource, pillarTarget)
}
