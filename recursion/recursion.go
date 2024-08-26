package recursion

func Recursion(i int) int {
	if i <= 2 {
		return i
	}
	return i * Recursion(i-1)
}
