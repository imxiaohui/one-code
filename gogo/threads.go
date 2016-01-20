package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Printf("thread runing!\n")
	time.Sleep(2 * time.Second)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i <= 10; i++ {
		go Count(lock)
	}
	fmt.Printf("thread starting!\n")
	go func() {
		for {
			lock.Lock()
			tmp := counter
			fmt.Println(tmp)
			time.Sleep(1 * time.Second)
			lock.Unlock()
			if tmp >= 10 {
				break
			}
		}
	}()
	time.Sleep(100 * time.Second)
}
