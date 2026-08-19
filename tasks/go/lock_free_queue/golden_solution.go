// Package queue implements a lock-free thread-safe concurrent queue in Go.
package queue

import (
	"sync/atomic"
	"unsafe"
)

type node struct {
	value unsafe.Pointer
	next  unsafe.Pointer
}

// LockFreeQueue represents a concurrent lock-free queue using CAS operations.
type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewLockFreeQueue() *LockFreeQueue {
	n := unsafe.Pointer(&node{})
	return &LockFreeQueue{
		head: n,
		tail: n,
	}
}

func (q *LockFreeQueue) Enqueue(v interface{}) {
	n := unsafe.Pointer(&node{value: unsafe.Pointer(&v)})
	for {
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*node)(tail).next)
		if tail == atomic.LoadPointer(&q.tail) {
			if next == nil {
				if atomic.CompareAndSwapPointer(&(*node)(tail).next, next, n) {
					atomic.CompareAndSwapPointer(&q.tail, tail, n)
					return
				}
			} else {
				atomic.CompareAndSwapPointer(&q.tail, tail, next)
			}
		}
	}
}
