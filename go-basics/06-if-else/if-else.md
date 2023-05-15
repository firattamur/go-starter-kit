## 06 - If/Else Statements in Go:

- Similar to other languages, Go has an `if` statement for conditional execution.

- Simple `if` statement:

  ```go
    if true {
        fmt.Println("Hello, Go!")
    }
  ```

- Simple `if/else` statement:

  ```go
    if true {
        fmt.Println("Hello, Go!")
    } else {
        fmt.Println("Goodbye, Go!")
    }
  ```

- `if/else if/else` statement:

  ```go
    if true {
        fmt.Println("Hello, Go!")
    } else if false {
        fmt.Println("Goodbye, Go!")
    } else {
        fmt.Println("See you later, Go!")
    }
  ```

- `if` statements can also be used to initialize variables:

  ```go
    if a := 42; a == 42 {
        fmt.Println("a: ", a)
    }
  ```

- `if` statements can also be used to initialize multiple variables:

  ```go
    if a, b := 42, 3.14; a == 42 {
        fmt.Println("a: ", a)
        fmt.Println("b: ", b)
    }
  ```

- Go does not have a ternary `?:` operator. You must use an `if` statement instead:

  ```go
    package main

    import "fmt"

    func main() {
        a := 42
        b := 3.14

        if a == 42 {
            fmt.Println("a: ", a)
        } else {
            fmt.Println("b: ", b)
        }
    }
  ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `if_else.go` which does the following:

1. Declare an `int` variable named `a` and assign the value 10 to it.

2. Declare an `int` variable named `b` and assign the value 20 to it.

3. Write an `if/else` statement that prints the larger of the two variables.

You can start with the following code skeleton:

```go
package main

import "fmt"

func main() {
    // Your code here
}
```

After you have finished writing your code, you can run it with the following command:

```bash
$ go run if_else.go
```

The expected output of your program should be:

```bash
20
```

This challenge will help you practice using if/else statements in Go. Good luck! 🍀
