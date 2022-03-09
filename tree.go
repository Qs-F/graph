package graph

type Node[T any] struct {
	Value    T
	Children []*Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}

func (node *Node[T]) Add(children ...*Node[T]) {
	node.Children = append(node.Children, children...)
}

func (node *Node[T]) DFS(f func(value T) bool) {
	var walk func(node *Node[T])
	walk = func(node *Node[T]) {
		if !f(node.Value) {
			return
		}
		if len(node.Children) == 0 {
			return
		}
		for _, child := range node.Children {
			walk(child)
		}
	}
	walk(node)
}
