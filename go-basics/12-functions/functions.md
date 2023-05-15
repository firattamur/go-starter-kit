## 12 - Functions in Go:

- Functions are declared using the `func` keyword:

  ```go
    func foo() {
        fmt.Println("Hello, Go!")
    }
  ```

- Functions can take arguments:

  ```go
    func foo(a int, b int) {
        fmt.Println(a + b)
    }
  ```

- Functions can take multiple arguments of the same type:

  ```go
    func foo(a, b int) int {
        return a + b
    }
  ```

- Functions can return values:

  ```go
    func foo(a int, b int) int {
        return a + b
    }
  ```

- Functions can return multiple values:

  ```go
    func foo(a int, b int) (int, int) {
        return a + b, a - b
    }
  ```

- Functions can return named values:

  ```go
    func foo(a int, b int) (c int, d int) {
        c = a + b
        d = a - b
        return // naked return statement with named values
    }

    func main() {
        c, d := foo(1, 2)
        fmt.Println(c, d)
    }
  ```

    - Output:
    
        ```bash
        3 -1
        ```

- Use `_` to ignore a value returned from a function:

  ```go
    func foo(a int, b int) (int, int) {
        return a + b, a - b
    }

    func main() {
        c, _ := foo(1, 2)
        fmt.Println(c)
    }
  ```

    - Output:
    
        ```bash
        3
        ```

- Functions can be variadic, meaning they take a variable number of arguments:

  ```go
    func foo(a ...int) {
        fmt.Println(a)
    }

    func main() {
        foo(1, 2, 3)
        foo(1, 2, 3, 4, 5)
    }
  ```

    - Output:
    
        ```bash
        [1 2 3]
        [1 2 3 4 5]
        ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `functions.go` which does the following:

1. Declare a function named `hello` that takes no arguments and prints "Hello, Go!".

2. Declare a function named `sum` that takes two `int` arguments and prints their sum.

3. Declare a function named `sumAndDiff` that takes two `int` arguments and returns their sum and difference.

4. Declare a function named `sumAndDiffv2` that takes two `int` arguments and returns their sum and difference as named values.

5. Declare a function named `sumAndDiffv3` that takes two `int` arguments and returns their sum and difference as named values. Use a naked return statement.

6. Declare a function name `multiSum` that takes a variadic number of `int` arguments and returns their sum.


You can start with the following code skeleton:

```go
package main

import "fmt"

// Your code here

func main() {
    
    hello()
    fmt.Println(sum(1, 2))
    fmt.Println(sumAndDiff(1, 2))
    fmt.Println(sumAndDiffv2(1, 2))
    fmt.Println(sumAndDiffv3(1, 2))
    fmt.Println(multiSum(1, 2, 3, 4, 5))

}
```

To run your program, use this command:

```bash
$ go run functions.go
```

The expected output of your program should be:

```bash
Hello, Go!
3
3 -1
3 -1
3 -1
15
```

This challenge will help you practice using the `func` keyword, declaring functions, and using variadic functions in Go. Good luck! 🍀