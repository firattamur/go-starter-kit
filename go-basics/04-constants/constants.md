## 04 - Constants

- Constants are declared like variables, but with the const keyword.

- Constants can be character, string, boolean, or numeric values.

- Constants cannot be declared using the := syntax.

  ```go
  const a := 42 // will raise the error: syntax error: unexpected :=, expecting =
  ```

- An untyped constant takes the type needed by its context.

- Let's look at some examples of declaring constants:

  ```go
  const a = 42
  const b = 3.14
  const c = "Hello World"
  const d = true
  ```

- Constants can also be declared using parentheses:

  ```go
    const (
        a = 42
        b = 3.14
        c = "Hello World"
        d = true
    )
  ```

- Constants can also be declared using the iota identifier:

  ```go
    const (
        a = iota
        b = iota
        c = iota
    )
  ```

  - iota starts at 0 and increments by one for each item in the list.

- Let's look at another example:

  ```go
  const (
      a = iota
      b
      c
  )
  ```

  - In this example, b and c will also be set to 1 and 2, respectively.

- Let's look a complete program that uses constants:

  ```go
  package main

  import "fmt"

  const (
      a = iota
      b
      c
  )

  const (
      d = iota
      e
      f
  )

  func main() {
      fmt.Println(a, b, c)
      fmt.Println(d, e, f)
  }
  ```

  - This program will print:

    ```bash
    0 1 2
    0 1 2
    ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `constants.go` which does the following:

1. Declare a set of constants for the first 5 prime numbers, using the `iota` identifier.

2. Print each of these constants.

You can start with the following code skeleton:

```go
package main

import "fmt"

// Your constants here

func main() {
// Your code here
}
```

To run your program, use this command:

```bash
$ go run constants.go
```

The expected output of your program should be:

```bash
2 3 5 7 11
```

This challenge will help you practice declaring constants and using the `iota` identifier in Go. Good luck! 🍀
