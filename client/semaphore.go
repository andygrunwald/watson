package client

import (
	"sync"
)

type semaphore chan bool

func NewSemaphore(n int) sync.Locker {
	s := make(semaphore, n)
	return s
}

func (s semaphore) Lock() {
	s <- true
}

func (s semaphore) Unlock() {
	<-s
}
