package main

import (
	"fmt"
	"flag"
	"time"
)


var period = flag.Duration("period", 1*time.Second, "sleep period")
var test = flag.String("test", "flag test", "test usage")

func main() {
    flag.Parse()
    fmt.Printf("Sleeping for %v...", *period)
    fmt.Println()
    fmt.Printf("custom for %v...", *test)
    time.Sleep(*period)
    fmt.Println()
}