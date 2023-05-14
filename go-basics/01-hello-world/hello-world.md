## 01 - Hello World in Go!

As is tradition, we'll start with a "Hello World" program.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

To run this program, we'll need to use the `go run` command.

```bash
$ go run hello-world.go

Hello World!
```

Let's break down the code above.

```go
package main
```

Every Go program starts with a package declaration. Packages are Go's way of organizing and reusing code. The package `main` is special. It defines a standalone executable program, not a library.

```go
func main() {
    // ...
}
```

Next we have the `main` function. It is the entry point for the executable program. It has no parameters and no return value. In Go, curly braces `{}` are always required, even if the body of the function is empty.

```go
fmt.Println("Hello World!")
```

Finally, we call a function called `Println` from the `fmt` package. Notice that we access the function using dot notation. This is called a qualified identifier in Go.
