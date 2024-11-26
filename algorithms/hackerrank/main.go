package hackerrank

func findChampion(n int, edges [][]int) int {

	strongerMap := make(map[int]int)

	for i := range edges {
		edge := edges[i]
		_, ok := strongerMap[edge[1]]
		if !ok {
			strongerMap[edge[1]] = 1
			continue
		}
		strongerMap[edge[1]]++
	}

	if len(strongerMap) != (n - 1) {
		return -1
	}

	var stronger int

	for i := 0; i < n; i++ {
		_, ok := strongerMap[i]
		if !ok {
			stronger = i
			break
		}
	}

	return stronger
}
