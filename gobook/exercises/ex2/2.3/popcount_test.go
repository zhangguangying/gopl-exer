package __3

import (
	"fmt"
	"gopl-exer/gobook/examples/ch2/popcount"
	__5 "gopl-exer/gobook/exercises/ex2/2.5"
	"testing"
	"time"
)

func TestPopCount(t *testing.T) {
	var i uint64 = 1024311212312312
	tabStart := time.Now()
	fmt.Println(popcount.PopCount(i))
	fmt.Println(time.Since(tabStart).Nanoseconds())

	cycleStart := time.Now()
	fmt.Println(PopCount(i))
	fmt.Println(time.Since(cycleStart).Nanoseconds())

	fmt.Println(__5.PopCount(i))
}

func TestUintTrans(t *testing.T) {
	for i := 0; i < 8; i++ {
		fmt.Println(uint(i))
	}
}