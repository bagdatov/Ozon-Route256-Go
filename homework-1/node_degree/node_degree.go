package nodedegree

import (
	"errors"
	"fmt"
)

var (
	errNotFound = errors.New("not found in the graph")
)

// Degree func
func Degree(nodes int, graph [][2]int, node int) (int, error) {
	if node > nodes {
		return 0, fmt.Errorf("node %d %w", node, errNotFound)
	}

	var deg int
	for _, edge := range graph {
		if node == edge[0] || node == edge[1] {
			deg++
		}
	}

	return deg, nil
}
