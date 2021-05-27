package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")

	time.Sleep(10 * time.Second)
}

func trace(msg string) string {
	log.Printf("enter %s", msg)
	return msg
}

func main() {
	bigSlowOperation()
}