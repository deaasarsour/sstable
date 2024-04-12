package awaitable

type Awaitable[R any] struct {
	callbackChan chan R
}

func (awaitable *Awaitable[R]) GetResult() R {
	return <-awaitable.callbackChan
}

func (awaitable *Awaitable[R]) SetResult(result R) {
	awaitable.callbackChan <- result
}

type AwaitableGroup[R any] struct {
	awaitables []*Awaitable[R]
}

func (awaitableGroup *AwaitableGroup[R]) SetResult(r R) {
	for _, awaitable := range awaitableGroup.awaitables {
		awaitable.SetResult(r)
	}
}

func NewAwaitable[R any]() *Awaitable[R] {
	return &Awaitable[R]{
		callbackChan: make(chan R, 1),
	}
}

func NewAwaitableGroup[R any](awaitables []*Awaitable[R]) *AwaitableGroup[R] {
	return &AwaitableGroup[R]{
		awaitables: awaitables,
	}
}
