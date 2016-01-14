package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	time.Sleep(10000)
	lock.Unlock()
	runtime.Gosched()
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i <= 10; i++ {
		go Count(lock)
	}
	go func() {
		for {
			lock.Lock()
			tmp := counter
			fmt.Println(tmp)
			time.Sleep(1000)
			lock.Unlock()
			if tmp >= 10 {
				break
			}
			runtime.Gosched()
		}
	}()
	time.Sleep(1000000)
}
