package proto

type ServerStream[T any] interface {
	Recv() (T, error)
	Ch() <-chan T
	Close() error
}

type serverStreamImpl[T any] struct {
	respCh  chan T
	closeCh chan struct{}
}

func (ss *serverStreamImpl[T]) Recv() (T, error) {
	return <-ss.respCh, nil
}

func (ss *serverStreamImpl[T]) Ch() <-chan T {
	return ss.respCh
}

func (ss *serverStreamImpl[T]) Close() error {
	close(ss.closeCh)
	return nil
}

type dummyStreamImpl[T any] struct {
	ch chan T
}

func NewDummyServerStream[T any]() ServerStream[T] {
	return &dummyStreamImpl[T]{make(chan T)}
}

func (ss *dummyStreamImpl[T]) Recv() (T, error) {
	return <-ss.ch, nil
}

func (ss *dummyStreamImpl[T]) Ch() <-chan T {
	return ss.ch
}

func (ss *dummyStreamImpl[T]) Close() error {
	return nil
}
