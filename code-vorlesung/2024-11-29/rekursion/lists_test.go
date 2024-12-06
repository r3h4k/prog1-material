package rekursion

import "fmt"

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

// Count2 erwartet eine Liste l und eine Zahl x.
// Zählt, wie oft x in l vorkommt.
func Count2(l []int, x int) int {
	if len(l) == 0 {
		return 0
	}

	r := Count2(l[1:], x)
	if l[0] == x {
		return 1 + r
	}
	return r
}

func ExampleCount() {
	fmt.Println(Count([]int{1, 2, 3, 4, 5, 3}, 3))
	fmt.Println(Count2([]int{1, 2, 3, 4, 5, 3}, 3))

	// Output:
	// 2
	// 2
}
