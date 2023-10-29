package graph

type AbstractGraph[T any] struct {
	Vertices []*Vertex[T]
}

type Vertex[T any] struct {
	Val   T
	Edges []*Edge[T]
}

type Edge[T any] struct {
	Weight     int
	DestVertex *Vertex[T]
}

func (g *AbstractGraph[T]) AddVertex(val T) error {
	newVertex := &Vertex[T]{Val: val, Edges: make([]*Edge[T], 0)}

	g.Vertices = append(g.Vertices, newVertex)

	return nil
}

func (g *AbstractGraph[T]) AddEdge(srcIndex, destIndex, weight int) error {
	newEdge := &Edge[T]{
		Weight:     weight,
		DestVertex: g.Vertices[destIndex],
	}

	g.Vertices[srcIndex].Edges = append(g.Vertices[srcIndex].Edges, newEdge)

	return nil
}
