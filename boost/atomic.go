package boost

import "sync/atomic"

type (
	// AtomicBool is a wrapper for atomic int32
	AtomicBool struct {
		value int32
	}
)

// Store atomicly store bool value
func (ab *AtomicBool) Store(b bool) {
	if b {
		atomic.StoreInt32(&ab.value, 1)
	} else {
		atomic.StoreInt32(&ab.value, 0)
	}
}

// Load atomicly load bool value
func (ab *AtomicBool) Load() bool {
	return atomic.LoadInt32(&ab.value) > 0
}
