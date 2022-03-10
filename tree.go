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
	visited := map[*Node[T]]struct{}{}
	var walk func(node *Node[T])
	walk = func(node *Node[T]) {
		if _, ok := visited[node]; ok {
			return
		}
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

func (node *Node[T]) BFS(f func(value T) bool) {
	q := NewQueue[*Node[T]](node)
	visited := map[*Node[T]]struct{}{node: struct{}{}}

	for q.Len() != 0 {
		v := q.Pop()
		if !f(v.Value) {
			return
		}
		for _, child := range v.Children {
			if _, ok := visited[child]; ok {
				continue
			}
			visited[child] = struct{}{}
			q.Push(child)
		}
	}
}
