package graph

// key is the address of value, value is the list of weights of edges
type AdjListGraph[T any] map[*T][]int

func (g *AdjListGraph[T]) AddVertex(value T) {
	(*g)[&value] = []int{}
}

//func (g *AdjListGraph[T]) AddEdge(src, dest, weight int) {
//	for key, val := range *g {
//		if key == src {
//
//		}
//	}
//}
