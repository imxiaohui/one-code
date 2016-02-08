package main

import "fmt"

func main() {

	c1 := make(chan string, 2)
	c1 <- "li"
	c1 <- "xu"
	c1 <- "12"
	c1 <- 1
	fmt.Printf("%s \n", <-c1)
	fmt.Printf("%s \n", <-c1)
}
