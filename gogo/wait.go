package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WaitObject struct {
	thread_num int
	msg        chan string
}

func createWait(thread_num int) *WaitObject {
	return &WaitObject{thread_num: thread_num, msg: make(chan string, thread_num)}
}

func (this *WaitObject) done() {
	this.msg <- "done"
}

func (this *WaitObject) wait() {
	for i := 0; i < this.thread_num; i++ {
		<-this.msg
	}
}

func main() {
	wait_object := createWait(10)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		go func() {
			rand_time := time.Duration((3 + rand.Intn(5)))
			time.Sleep(rand_time * time.Second)
			fmt.Println("running !\n")
			wait_object.done()
		}()
	}
	wait_object.wait()
}
