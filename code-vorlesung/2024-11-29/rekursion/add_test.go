package rekursion

import "fmt"

func Add1(x, y int) int {
	if y == 0 {
		return x
	}

	a := Add1(x, y-1)
	r := a + 1
	return r
}

func Add2(x, y int) int {
	if y == 0 {
		return x
	}

	return Add2(x+1, y-1)
}

func Add3(x, y int) int {
	for y != 0 {
		x++
		y--
	}

	return x
}

func ExampleAdd1() {
	r := Add1(2, 3)
	fmt.Println(r)

	// Output:
}
