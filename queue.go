package graph

type Queue[T any] []T

func NewQueue[T any](values ...T) *Queue[T] {
	ret := Queue[T](values)
	return &ret
}

func (queue *Queue[T]) Push(values ...T) {
	*queue = append(*queue, values...)
}

func (queue *Queue[T]) Peek() T {
	return (*queue)[0]
}

func (queue *Queue[T]) Pop() T {
	ret := (*queue)[0]
	*queue = (*queue)[1:]
	return ret
}

func (queue *Queue[T]) Len() int {
	return len(*queue)
}
