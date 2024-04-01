package util

type Future[R any] struct {
	returnChan chan R
}

func (future *Future[R]) GetResult() R {
	return <-future.returnChan
}

func (future *Future[R]) SetResult(r R) {
	future.returnChan <- r
}

type FutureGroup[R any] struct {
	futures []*Future[R]
}

func (futureGroup *FutureGroup[R]) SetResult(r R) {
	for _, future := range futureGroup.futures {
		future.SetResult(r)
	}
}

func NewFuture[R any]() *Future[R] {
	return &Future[R]{
		returnChan: make(chan R, 1),
	}
}

func NewFutureGroup[R any](futures []*Future[R]) *FutureGroup[R] {
	return &FutureGroup[R]{
		futures: futures,
	}
}
