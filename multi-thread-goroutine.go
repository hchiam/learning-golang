// concurrency (not necessarily parallelism)
// learn more about concurrency, channels, and application example exercises, at https://go.dev/tour/concurrency/1
// you can format and run code for free at https://go.dev/play/

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("1")
	say("hello world")
}
