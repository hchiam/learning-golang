// concurrency (not necessarily parallelism)
// learn more about concurrency, channels, and application example exercises, at https://go.dev/tour/concurrency/1

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
