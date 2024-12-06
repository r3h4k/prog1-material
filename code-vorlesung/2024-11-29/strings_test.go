package main

import "fmt"

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
	fmt.Println(ContainsChar("abcd", 'c'))
	fmt.Println(ContainsChar("abcd", 'd'))
	fmt.Println(ContainsChar("", 'a'))

	// Output:
	// true
	// true
	// true
	// true
	// false

}
