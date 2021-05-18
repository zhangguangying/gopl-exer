package basename2

import (
	"fmt"
	"testing"
)

func TestBasename(t *testing.T) {
	s1 := "a.go"
	s2 := "a/b/c.go"
	s3 := "a/b.c.go"
	fmt.Println(Basename(s1))
	fmt.Println(Basename(s2))
	fmt.Println(Basename(s3))
}
