package popcount

import (
	"fmt"
	"testing"
)

func TestPopCount(t *testing.T) {
	fmt.Println(PopCount(63))
	var s uint64 = 512
	fmt.Printf("%b\n", s)
	fmt.Printf("%b\n", byte(s))
}