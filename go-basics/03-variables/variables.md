## 03 - Variables in Go:

- Variables are used to store values in memory.

- Go is a statically typed language, which means that variables always have a specific type and that type cannot change.

- Go is also a strongly typed language, which means that variables can only be coerced to another type with an explicit conversion.

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

- Short variable declarations are also possible:

  ```go
    a := 42
  ```

- We can also declare multiple variables at once:

  ```go
    var a, b int = 1, 2
  ```

- Go also supports parallel assignment:

  ```go
    a, b := 1, 2
  ```

- Go also supports assigning multiple variables at once:

  ```go
    a, b := 1, 2
  ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `variables.go` which does the following:

1. Declare an `int` variable named `a` and assign the value 10 to it.
2. Declare an `int` variable named `b` and assign the value 20 to it.
3. Use parallel assignment to swap the values of `a` and `b`.
4. Print the values of `a` and `b` after the swap.

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
$ go run variables.go
```

The expected output of your program should be:

```bash
20 10
```

This challenge will help you practice declaring variables, assigning values, and using parallel assignment in Go. Good luck! 🍀
