# Learn Golang

Just one of the things I'm learning: <https://github.com/hchiam/learning>

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

Run binary/executable of `hello.go`:

```bash
./hello
```

## To install Golang linter

```bash
go get -u golang.org/x/lint/golint
```
