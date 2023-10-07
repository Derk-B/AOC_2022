package math

// A function to replace the % operator.
// Go % is a remainder operator and not a modulo.
func Mod(a int, n int) int {
	return ((a % n) + n) % n
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
	
