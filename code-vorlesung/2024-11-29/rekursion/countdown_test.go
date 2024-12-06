package rekursion

import "fmt"

func Countdown(n int) {
	if n <= 0 {
		return
	}

	fmt.Println(n)
	Countdown(n - 1)
	fmt.Println(n)
}

func ExampleCountdown() {
	Countdown(3)

	// Output:
}
