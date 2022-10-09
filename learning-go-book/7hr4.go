package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg = sync.WaitGroup{}

var counter = 0
var m = sync.RWMutex{}

func main () {
//	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		wg.Add(2)
	m.RLock()
		go sayHello()
	m.Lock()
		go increment()
	}
	wg.Wait()
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
