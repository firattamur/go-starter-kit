## 02 - Data Types in Go:

- Go is a statically typed language. This means that variables always have a specific type and that type cannot change.

- Go is also a strongly typed language, which means that variables can only be coerced to another type with an explicit conversion.

- Go has the following built-in types:

  - `bool`
  - `string`
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
  - `byte` (alias for `uint8`)
  - `rune` (alias for `int32`, represents a Unicode code point)
  - `float32`, `float64`
  - `complex64`, `complex128`

- Go does not have a `char` type. Instead, `rune` is used to represent a Unicode code point.

- To declare a variable in Go, we use the `var` keyword:

  ```go
  var a int
  ```

- The type comes after the variable name.

- We can also initialize the variable with a value:

  ```go
  var a int = 42
  ```

- Go can infer the type of initialized variables:

  ```go
  var a = 42
  ```

- Let's look at some examples of declaring variables with different types:

  ```go
  var a bool = true
  var b string = "Hello World"
  var c int = 42
  var d float64 = 3.14
  var e complex128 = 3 + 5i
  var f byte = 'a'
  var g rune = '☺'
  ```

- We can also declare multiple variables in one statement:

  ```go
    var a, b, c int = 1, 2, 3
  ```

- Go also has short variable declarations using the `:=` operator:

  ```go
    a := 42
  ```

- Short variable declarations can only be used inside functions.

## 🚀 Coding Challenge:

Your task is to write a Go program named `types.go` which does the following:

1. Declare a `string` variable named `hello` and assign the value "Hello, Go!" to it.
2. Declare an `int` variable named `year` and assign the current year to it.
3. Print both variables using the `fmt.Println()` function.

You can start with the following code skeleton:

```go
package main

import "fmt"

func main() {
// Your code here
}
```

To run your program, use this command:

```bash
$ go run types.go
```

The expected output of your program should be:

```bash
Hello, Go! 2023
```

This challenge will help you practice declaring variables, assigning values, and printing values in Go. Good luck! 🍀
