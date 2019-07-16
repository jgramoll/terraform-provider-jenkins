package multilock

import (
	"sync"
	"sync/atomic"
)

// A multi lock type
type lock struct {
	inUse sync.Map
	pool  *sync.Pool
}

// NewBasicMultiLock create a new basic lock for multiple keys
func NewBasicMultiLock() MultiLock {
	return &lock{
		pool: &sync.Pool{
			New: func() interface{} {
				return &sync.RWMutex{}
			},
		},
	}
}

func (l *lock) Lock(key interface{}) {
	m := l.getLocker(key)
	atomic.AddInt64(&m.counter, 1)
	m.lock.Lock()
}

func (l *lock) RLock(key interface{}) {
	m := l.getLocker(key)
	atomic.AddInt64(&m.counter, 1)
	m.lock.RLock()
}

func (l *lock) Unlock(key interface{}) {
	m := l.getLocker(key)
	m.lock.Unlock()
	l.putBackInPool(key, m)
}

func (l *lock) RUnlock(key interface{}) {
	m := l.getLocker(key)
	m.lock.RUnlock()
	l.putBackInPool(key, m)
}

func (l *lock) putBackInPool(key interface{}, m *refCounter) {
	atomic.AddInt64(&m.counter, -1)
	if m.counter <= 0 {
		l.pool.Put(m.lock)
		l.inUse.Delete(key)
	}
}

func (l *lock) getLocker(key interface{}) *refCounter {
	res, _ := l.inUse.LoadOrStore(key, &refCounter{
		counter: 0,
		lock:    l.pool.Get().(*sync.RWMutex),
	})

	return res.(*refCounter)
}
