package multilock

import (
	"sync"
)

type refCounter struct {
	counter int64
	lock    *sync.RWMutex
}