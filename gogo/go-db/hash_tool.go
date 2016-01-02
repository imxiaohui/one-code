package godb

import (
	"fmt"
	"hash"
)

func hashInt(value int, table_size int) int {
	return value % table_size
}

func hashString(value []byte, table_size int) int {
	return value
}
