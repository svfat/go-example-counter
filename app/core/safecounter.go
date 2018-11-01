package core

import "sync"

type Counter struct {
	value     uint64
	countLock sync.Mutex
}

// Increment counter
func (cnt *Counter) Inc() {
	cnt.countLock.Lock()
	cnt.value++
	cnt.countLock.Unlock()
}

// Set counter
func (cnt *Counter) Set(value uint64) {
	cnt.countLock.Lock()
	cnt.value = value
	cnt.countLock.Unlock()
}

// Get the value of counter
func (cnt *Counter) Value() uint64 {
	cnt.countLock.Lock()
	defer cnt.countLock.Unlock()
	return cnt.value
}
