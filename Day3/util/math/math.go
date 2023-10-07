package math

// A function to replace the % operator.
// Go % is a remainder operator and not a modulo.
func Mod(a int, n int) int {
	return ((a % n) + n) % n
}

