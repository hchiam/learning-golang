# Learning Golang

Just one of the things I'm learning: <https://github.com/hchiam/learning>

You can format and run go code for free in the browser: <https://go.dev/play/>

Learn more about concurrency, channels, and application example exercises, at <https://go.dev/tour/concurrency/1>

Tutorial from Google and Codecademy: <https://www.codecademy.com/learn/learn-go>

## [Basic local setup](https://www.codecademy.com/articles/setting-up-go-locally)

Download Golang: <https://golang.org/dl>

**Aside:** You might also want to set up your `.bash_profile`:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Build binary/executable of `hello.go`:

```bash
go build hello.go
```

Run binary/executable named simply `hello`:

```bash
./hello
```

## Compile and execute in one step*

(*BUT does NOT create executable `hello`):

```bash
go run hello.go
```

## To install Golang linter

```bash
go get -u golang.org/x/lint/golint
```

## To get documentation on something

Example: to make the CLI show documentation on the `Println` function of `fmt`:

```bash
go doc fmt.Println
```

## Miscellaneous notes

- `:=` = declare variable and assign value (type can be inferenced), e.g. `var x int := 1`
- `=` = re-assign (variable must already exist in scope), e.g. `x = 2`
- `go something()` = run `something()` concurrently (in a "goroutine")
- to combine results from multiple goroutines, you can use channels, which are blocking until the other side is ready when they send/receive (allowing goroutines to synchronize without explicit locks or condition variables):
  - e.g.:

        ```go
        package main

        import "fmt"

        func sum(s []int, c chan int) {
            sum := 0
            for _, v := range s {
                sum += v
            }
            c <- sum // send sum to channel c
        }

        func main() {
            s := []int{7, 2, 8, -9, 4, 0}

            bufferSize := 1 // optional. "Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty."
            c := make(chan int, bufferSize) // create and use the same channel c for the goroutines:
            go sum(s[:len(s)/2], c) // process 1st half concurrently
            go sum(s[len(s)/2:], c) // process 2nd half concurrently
            x, y := <-c, <-c // receive from channel c

            fmt.Println(x, y, x+y)
        }
        ```
    - not quite sure what "Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty." means in <https://go.dev/tour/concurrency/3>
- `value, didChannelClose := <-c`
  - "Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic." <https://go.dev/tour/concurrency/4>
    - i.e., the goroutine function being called should `close(c)`
  - "Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop." <https://go.dev/tour/concurrency/4>
- select statement `for { select { case c <- x: ... case <-quit: ... return } }` "lets a goroutine wait on multiple communication operations" <https://go.dev/tour/concurrency/5> and blocks until one case can run, and if multiple are ready it chooses one at random to run.
