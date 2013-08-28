// see: https://raw.github.com/abiosoft/semaphore/master/semaphore.go
package channels

import (
	"sync"
	"time"
)

type Semaphore struct {
	permits int
	avail   int
	channel chan int
	aMutex  *sync.Mutex // acquire Mutex
	rMutex  *sync.Mutex // release Mutex
}

func NewSemaphore(permits int) *Semaphore {
	if permits < 1 {
		panic("Invalid number of permits. Less than 1")
	}
	return &Semaphore{
		permits,
		permits, // same available as allocated
		make(chan int, permits),
		&sync.Mutex{}, // create a new mutex, return the address
		&sync.Mutex{},
	}
}

//Acquire one permit, if its not available the goroutine will block till its available
func (s *Semaphore) Acquire() {
	s.AcquireMany(1) // just acquire 1
}

//Similar to Acquire() but for many permits
func (s *Semaphore) AcquireMany(n int) {
	if n > s.permits {
		panic("To many requested permits")
	}
	s.aMutex.Lock() // acquire a lock
	s.avail -= n    // decrement the number of available by n
	for ; n > 0; n-- {
		s.channel <- 1 // pump in a 1 into the channel
	}
	s.avail += n      // increment the number available
	s.aMutex.Unlock() // unlock the acquire lock
}

//Similar to AcquireMany() but cancels if duration elapse before getting the permits.
//Returns true if successful and false if timeout occurs.
func (s *Semaphore) AcquireWithin(n int, d time.Duration) bool {
	timeout := make(chan bool, 1)
	cancel := make(chan bool, 1)
	go func() {
		time.Sleep(d)
		timeout <- true
	}()
	go func() {
		s.AcquireMany(n)
		timeout <- false
		if <-cancel {
			s.ReleaseMany(n)
		}
	}()
	if <-timeout {
		cancel <- true
		return false
	}
	cancel <- false
	return true
}

//Release one permit
func (s *Semaphore) Release() {
	s.rMutex.Lock()
	<-s.channel
	s.avail++
	s.rMutex.Unlock()
}

//Release many permits
func (s *Semaphore) ReleaseMany(n int) {
	if n > s.permits {
		panic("Too many requested releases")
	}
	for ; n > 0; n-- {
		s.Release()
	}
}

//Number of available unacquired permits
func (s *Semaphore) AvailablePermits() int {
	if s.avail < 0 {
		return 0
	}
	return s.avail
}

//Acquire all available permits and return the number of permits acquired
func (s *Semaphore) DrainPermits() int {
	n := s.AvailablePermits()
	if n > 0 {
		s.AcquireMany(n)
	}
	return n
}
