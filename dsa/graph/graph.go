package graph

type Graph[T any] interface {
	AddVertex(value T) error
	AddEdge(weight int) error
	PrintVertex()
}
