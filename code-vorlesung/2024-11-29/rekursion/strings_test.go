package rekursion

import "fmt"

// ContainsChar erwartet einen String s und ein Zeichen c.
// Liefer true, falls c in s vorkommt.
func ContainsChar(s string, c byte) bool {
	if s == "" {
		return false
	}

	if s[0] == c {
		return true
	}
	return ContainsChar(s[1:], c)
}

func ExampleContainsChar() {
	fmt.Println(ContainsChar("abc", 'a'))
	fmt.Println(ContainsChar("abc", 'b'))
	fmt.Println(ContainsChar("abc", 'c'))
	fmt.Println(ContainsChar("abc", 'd'))

	// Output:
	// true
	// true
	// true
	// false
}
