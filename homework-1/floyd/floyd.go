package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	if rows <= 0 {
		return nil
	}

	var (
		triangle = make([][]int, rows)
		num      = 1
	)

	for y := 0; y < rows; y++ {

		triangle[y] = make([]int, y+1)

		for x := 0; x <= y; x++ {
			triangle[y][x] = num
			num++
		}
	}

	return triangle
}
