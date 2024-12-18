package ackermann

import "fmt"

func Ackermann(m, n int) int {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return Ackermann(m-1, 1)
	} else {
		return Ackermann(m-1, Ackermann(m, n-1))
	}
}

func ExampleAckermann() {
	fmt.Println(Ackermann(2, 0))

	// Output:
}
