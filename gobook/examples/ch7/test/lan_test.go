package test

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestType(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Println(f)
}

func TestType2(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	fmt.Printf("%T\n", w)
}

func TestType3(t *testing.T) {
	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	fmt.Println(os.IsNotExist(err))
}