package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		value, err := <-c
		if err {
			fmt.Println("cool", value)
		}
	}()
	go func() {

		for {
			value, err := <-c
			if err {
				fmt.Println("cool", value)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(13 * time.Second)
	c <- 1
	time.Sleep(5 * time.Second)
}
