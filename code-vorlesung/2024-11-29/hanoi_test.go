package main

import "fmt"

func Hanoi(h int, s, m, z string) {
	if h == 0 {
		return
	}
	Hanoi(h-1, s, z, m)
	fmt.Printf("%s --> %s\n", s, z)
	Hanoi(h-1, m, s, z)
}
func ExampleHanoi() {
	Hanoi(3, "A", "B", "C")

	// Output:
}
