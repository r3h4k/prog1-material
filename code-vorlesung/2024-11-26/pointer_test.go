package pointer

import "fmt"

func Example_pointer() {
	var x int

	// Die Adresse von x in einem int-Pointer px speichern
	// Pointer darf nicht im nachhinein geändert werden
	var px *int = &x

	// x verändern
	x = 42

	// Speicheradresse von x ausgeben
	fmt.Println(px)

	// den Wert hiner px verändern.
	*px = 23

	// Den Wert von x ausgeben, aber durch den Pointer hindurch.
	// * folge diesen Pointer
	fmt.Println(x)

	// Output:
}

// addiert 2 auf den Wert von x
func Add2(x *int) {
	*x = *x + 2
}

// Multipliziert den Wert von x mit 2
func Mult2(x *int) {
	*x = *x * 2
}

func Example_call_by_ref() {
	x := 40
	y := 17
	Add2(&x)
	Mult2(&y)

	fmt.Println(x)

	// Output:
	// 42
}
