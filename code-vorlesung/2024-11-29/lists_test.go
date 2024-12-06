package main

// Count erwartet eine Liste l und eine Zahl x.
// Zählt, wie oft x in l vorkommt.
func Count(l []int, x int) int {
	if len(l) == 0 {
		return 0
	}

	if l[0] == x {
		return 1 + Count(l[1:], x)
	}
	return Count(l[1:], x)
}
