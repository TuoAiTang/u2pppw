package main

import (
	"fmt"
	"sync"
)

type AtomicInteger struct {
	val int
	lock sync.RWMutex
}

func (ai *AtomicInteger) Add(i int) {
	ai.lock.Lock()
	defer ai.lock.Unlock()
	defer ai.lock.Unlock()
	ai.lock.Lock()
	ai.val ++
}

func (ai *AtomicInteger) Get() int {
	return ai.val
}

func main() {
	var ai AtomicInteger

	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(integer *AtomicInteger) {
			integer.Add(1)
			wg.Done()
		}(&ai)
	}

	wg.Wait()
	fmt.Println(ai.Get())
}
