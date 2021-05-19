package __1

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestBitDiff(t *testing.T)  {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(bitDiff(c1, c2))
}
