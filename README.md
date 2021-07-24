![Go logo](Go-Logo_LightBlue.png "Go logo")

This is a basic Go package created to learn the language.

## Build & run

Go is a compiled language. To run a Go program we need to build it first.

### Build & run with `go install`

```zsh
$ go install
```

The program will be built into `$GOPATH/bin`. Look for an executable called `go-hello`.

### Build & run with `go run`

```zsh
$ go run hello.go
```

The program will be built and run immediately in the directory. This is good for quick testing of the program.
