package dijkstra

type robot struct {
	path []string
	pathLength float64
}

type dest struct {
	source string
	dest string
	pathLength float64
}