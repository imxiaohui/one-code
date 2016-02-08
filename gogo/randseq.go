package main

import (
	"fmt"
	"math/rand"
	"time"
)

const RAND_SEQ = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
const RAND_LEN = len(RAND_SEQ)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func rand_seq(n int) string {
	seq := make([]rune, n)
	for i := 0; i < n; i++ {
		seq[i] = RAND_SEQ[rand.Intn(RAND_LEN)]
	}
	return string(seq)
}

func main() {

	init()
	rand_seq(5)
}
