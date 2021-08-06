package main

import (
	"time"
	"log"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10*time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("Enter %s", msg)
	return func() {
		log.Printf("Exit %s (%s)", msg, time.Since(start))
	}
}