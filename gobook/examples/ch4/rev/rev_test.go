package rev

import (
	"fmt"
	"testing"
)

func TestRev(t *testing.T) {
	s := []int{1,2,3,4,5}
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}

func TestTrans(t *testing.T)  {
	s := []int{1,2,3,4,5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s[:])
	fmt.Println(s)
}
