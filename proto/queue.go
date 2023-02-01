package proto

type Queue[T any] struct {
	elems []T
}

func (q *Queue[T]) Push(t T) { q.elems = append(q.elems, t) }

func (q *Queue[T]) Pop() T {
	front := q.elems[0]
	q.elems = q.elems[1:]
	return front
}

func (q *Queue[T]) Len() int    { return len(q.elems) }
func (q *Queue[T]) Empty() bool { return len(q.elems) == 0 }
