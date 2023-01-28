package coins

func Piles(n int) int {
	return solve(n, 1)
}

func solve(n int, min int) int {
	if n == 0 || n == min {
		return 1
	}

	if n < min {
		return 0
	}

	var res int

	for i := min; i <= n; i++ {
		res += solve(n-i, i)
	}

	return res
}
