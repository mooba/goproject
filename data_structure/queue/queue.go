// author pengchengbai@shopee.com
// date 2021/2/2

package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type waiters []*sema

func (w *waiters) get() *sema {
	if len(*w) == 0 {
		return nil
	}

	sema := (*w)[0]
	copy((*w)[0:], (*w)[1:])
	(*w)[len(*w)-1] = nil // or the zero value of T
	*w = (*w)[:len(*w)-1]
	return sema
}

func (w *waiters) put(sema *sema) {
	*w = append(*w, sema)
}

func (w *waiters) remove(sema *sema) {
	if len(*w) == 0 {
		return
	}
	// build new slice, copy all except sema
	ws := *w
	newWs := make(waiters, 0, len(*w))
	for i := range ws {
		if ws[i] != sema {
			newWs = append(newWs, ws[i])
		}
	}
	*w = newWs
}

type items []interface{}

func (items *items) get(number int64) []interface{} {
	returnItems := make([]interface{}, 0, number)
	index := int64(0)
	for i := int64(0); i < number; i++ {
		if i >= int64(len(*items)) {
			break
		}

		returnItems = append(returnItems, (*items)[i])
		(*items)[i] = nil
		index++
	}

	*items = (*items)[index:]
	return returnItems
}

func (items *items) peek() (interface{}, bool) {
	length := len(*items)

	if length == 0 {
		return nil, false
	}

	return (*items)[0], true
}

func (items *items) getUntil(checker func(item interface{}) bool) []interface{} {
	length := len(*items)

	if len(*items) == 0 {
		// returning nil here actually wraps that nil in a list
		// of interfaces... thanks go
		return []interface{}{}
	}

	returnItems := make([]interface{}, 0, length)
	index := -1
	for i, item := range *items {
		if !checker(item) {
			break
		}

		returnItems = append(returnItems, item)
		index = i
		(*items)[i] = nil // prevent memory leak
	}

	*items = (*items)[index+1:]
	return returnItems
}

type sema struct {
	ready    chan bool
	response *sync.WaitGroup
}

func newSema() *sema {
	return &sema{
		ready:    make(chan bool, 1),
		response: &sync.WaitGroup{},
	}
}

// Queue is the struct responsible for tracking the state
// of the queue.
type Queue struct {
	waiters  waiters
	items    items
	lock     sync.Mutex
	disposed bool
}

// Put will add the specified items to the queue.
func (q *Queue) Put(items ...interface{}) error {
	if len(items) == 0 {
		return nil
	}

	q.lock.Lock()

	if q.disposed {
		q.lock.Unlock()
		return ErrDisposed
	}

	q.items = append(q.items, items...)
	for {
		sema := q.waiters.get()
		if sema == nil {
			break
		}
		sema.response.Add(1)
		select {
		case sema.ready <- true:
			sema.response.Wait()
		default:
			// This semaphore timed out.
		}
		if len(q.items) == 0 {
			break
		}
	}

	q.lock.Unlock()
	return nil
}

// Get retrieves items from the queue.  If there are some items in the
// queue, get will return a number UP TO the number passed in as a
// parameter.  If no items are in the queue, this method will pause
// until items are added to the queue.
func (q *Queue) Get(number int64) ([]interface{}, error) {
	return q.Poll(number, 0)
}

// Poll retrieves items from the queue.  If there are some items in the queue,
// Poll will return a number UP TO the number passed in as a parameter.  If no
// items are in the queue, this method will pause until items are added to the
// queue or the provided timeout is reached.  A non-positive timeout will block
// until items are added.  If a timeout occurs, ErrTimeout is returned.
func (q *Queue) Poll(number int64, timeout time.Duration) ([]interface{}, error) {
	if number < 1 {
		// thanks again go
		return []interface{}{}, nil
	}

	q.lock.Lock()

	if q.disposed {
		q.lock.Unlock()
		return nil, ErrDisposed
	}

	var items []interface{}

	if len(q.items) == 0 {
		sema := newSema()
		q.waiters.put(sema)
		q.lock.Unlock()

		var timeoutC <-chan time.Time
		if timeout > 0 {
			timeoutC = time.After(timeout)
		}
		select {
		case <-sema.ready:
			// we are now inside the put's lock
			if q.disposed {
				return nil, ErrDisposed
			}
			items = q.items.get(number)
			sema.response.Done()
			return items, nil
		case <-timeoutC:
			// cleanup the sema that was added to waiters
			select {
			case sema.ready <- true:
				// we called this before Put() could
				// Remove sema from waiters.
				q.lock.Lock()
				q.waiters.remove(sema)
				q.lock.Unlock()
			default:
				// Put() got it already, we need to call Done() so Put() can move on
				sema.response.Done()
			}
			return nil, ErrTimeout
		}
	}

	items = q.items.get(number)
	q.lock.Unlock()
	return items, nil
}

// Peek returns a the first item in the queue by value
// without modifying the queue.
func (q *Queue) Peek() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.disposed {
		return nil, ErrDisposed
	}

	peekItem, ok := q.items.peek()
	if !ok {
		return nil, ErrEmptyQueue
	}

	return peekItem, nil
}

// TakeUntil takes a function and returns a list of items that
// match the checker until the checker returns false.  This does not
// wait if there are no items in the queue.
func (q *Queue) TakeUntil(checker func(item interface{}) bool) ([]interface{}, error) {
	if checker == nil {
		return nil, nil
	}

	q.lock.Lock()

	if q.disposed {
		q.lock.Unlock()
		return nil, ErrDisposed
	}

	result := q.items.getUntil(checker)
	q.lock.Unlock()
	return result, nil
}

// Empty returns a bool indicating if this bool is empty.
func (q *Queue) Empty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items) == 0
}

// Len returns the number of items in this queue.
func (q *Queue) Len() int64 {
	q.lock.Lock()
	defer q.lock.Unlock()

	return int64(len(q.items))
}

// Disposed returns a bool indicating if this queue
// has had disposed called on it.
func (q *Queue) Disposed() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.disposed
}

// Dispose will dispose of this queue and returns
// the items disposed. Any subsequent calls to Get
// or Put will return an error.
func (q *Queue) Dispose() []interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.disposed = true
	for _, waiter := range q.waiters {
		waiter.response.Add(1)
		select {
		case waiter.ready <- true:
			// release Poll immediately
		default:
			// ignore if it's a timeout or in the get
		}
	}

	disposedItems := q.items

	q.items = nil
	q.waiters = nil

	return disposedItems
}

// New is a constructor for a new threadsafe queue.
func New(hint int64) *Queue {
	return &Queue{
		items: make([]interface{}, 0, hint),
	}
}

// ExecuteInParallel will (in parallel) call the provided function
// with each item in the queue until the queue is exhausted.  When the queue
// is exhausted execution is complete and all goroutines will be killed.
// This means that the queue will be disposed so cannot be used again.
func ExecuteInParallel(q *Queue, fn func(interface{})) {
	if q == nil {
		return
	}

	q.lock.Lock() // so no one touches anything in the middle
	// of this process
	todo, done := uint64(len(q.items)), int64(-1)
	// this is important or we might face an infinite loop
	if todo == 0 {
		return
	}

	numCPU := 1
	if runtime.NumCPU() > 1 {
		numCPU = runtime.NumCPU() - 1
	}

	var wg sync.WaitGroup
	wg.Add(numCPU)
	items := q.items

	for i := 0; i < numCPU; i++ {
		go func() {
			for {
				index := atomic.AddInt64(&done, 1)
				if index >= int64(todo) {
					wg.Done()
					break
				}

				fn(items[index])
				items[index] = 0
			}
		}()
	}
	wg.Wait()
	q.lock.Unlock()
	q.Dispose()
}
