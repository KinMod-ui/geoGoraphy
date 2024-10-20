package main


func CountNodes(start , end vertex) int {
	cnt := 0
	for i := start.x; i <= end.x; i++ {
		for j := start.y; j <= end.y; j++ {
			cnt += grid[j][i]
		}
	}
	return cnt
}
