package gonative

func fibtc(n, a, b int) int {
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else {
		return fibtc(n-1, b, a+b)
	}
}
