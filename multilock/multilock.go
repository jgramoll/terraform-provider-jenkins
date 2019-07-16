package multilock

// MultiLock is the main interface for lock base on key
type MultiLock interface {
	// Lock base on the key
	Lock(interface{})

	// RLock lock the rw for reading
	RLock(interface{})

	// Unlock the key
	Unlock(interface{})

	// RUnlock the the read lock
	RUnlock(interface{})
}
