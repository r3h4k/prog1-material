package hanoi

import "fmt"

// Move führt eine Scheibenbewegung von s nach z aus.
// In diesem Fall bedeutet das, dass die Bewegung auf der Konsole ausgegeben wird.
func Move(s, z string) {
	fmt.Printf("Bewege Scheibe von %s nach %s.\n", s, z)
}

/* Die folgenden Funktionen lösen das Problem jeweils für eine bestimmte Höhe.
Jede Funktion erwartet als Parameter die Name der drei Türme in der
Reihenfolge "start", "mitte", "ziel" (s,m,z).
*/

// Hanoi1 löst das Türme-von-Hanoi-Problem für Höhe 1.
func Hanoi1(s, m, z string) {
	Move(s, z)
}

// Hanoi2 löst das Türme-von-Hanoi-Problem für Höhe 2.
func Hanoi2(s, m, z string) {
	Hanoi1(s, z, m)
	Move(s, z)
	Hanoi1(m, s, z)
}

// Hanoi3 löst das Türme-von-Hanoi-Problem für Höhe 3.
func Hanoi3(s, m, z string) {
	Hanoi2(s, z, m)
	Move(s, z)
	Hanoi2(m, s, z)
}

// Hanoi4 löst das Türme-von-Hanoi-Problem für Höhe 4.
func Hanoi4(s, m, z string) {
	Hanoi3(s, z, m)
	Move(s, z)
	Hanoi3(m, s, z)
}

// Hanoi5 löst das Türme-von-Hanoi-Problem für Höhe 5.
func Hanoi5(s, m, z string) {
	Hanoi4(s, z, m)
	Move(s, z)
	Hanoi4(m, s, z)
}

// Hanoi6 löst das Türme-von-Hanoi-Problem für Höhe 6.
func Hanoi6(s, m, z string) {
	Hanoi5(s, z, m)
	Move(s, z)
	Hanoi5(m, s, z)
}

/* Die folgende Funktion löst das Problem für eine beliebige Höhe.
   Sie erwartet als Parameter die Höhe und die Nummern der drei Türme in der
   Reihenfolge "von", "über", "nach".
   Es fällt auf, dass die Funktion sich selbst aufruft, ansonsten aber
   das gleiche Muster wie die Funktionen für die einzelnen Höhen aufweist.
*/

// Hanoi löst das Türme-von-Hanoi-Problem für eine beliebige Höhe h.
func Hanoi(h int, s, m, z string) {
	if h == 1 {
		Move(s, z)
	} else {
		Hanoi(h-1, s, z, m)
		Move(s, z)
		Hanoi(h-1, m, s, z)
	}
}
