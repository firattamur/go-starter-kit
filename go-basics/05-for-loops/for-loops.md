## 05 - For Loops in Go

- Go has only one looping construct, the `for` loop.

- The basic `for` loop has three components separated by semicolons:

  ```go
  for initialization; condition; post {
    // zero or more statements
  }
  ```

  - The initialization statement is executed before the first iteration.

  - The condition expression is evaluated before every iteration.

  - The post statement is executed at the end of every iteration.

- Let's look at a simple example:

  ```go
    package main

    import "fmt"

    func main() {
        for i := 0; i < 5; i++ {
            fmt.Println(i)
        }
    }

  ```

  - This program will print:

    ```bash
    0
    1
    2
    3
    4
    ```

- The initialization and post statements are optional:

  ```go
    package main

    import "fmt"

    func main() {
        i := 0
        for i < 5 {
            fmt.Println(i)
            i++
        }
    }
  ```

  - This program will print the same output as the previous example.

- If you omit the loop condition, it loops forever:

  ```go
    package main

    import "fmt"

    func main() {
        i := 0
        for {
            fmt.Println(i)
            i++
        }
    }
  ```

  - This program will print forever.

- You can also use the `break` and `continue` statements inside a `for` loop:

  ```go
    package main

    import "fmt"

    func main() {
        for i := 0; i < 5; i++ {
            if i == 3 {
            continue
            }
            fmt.Println(i)
        }
    }
  ```

  - This program will print:

    ```bash
    0
    1
    2
    4
    ```

- To break out of a loop, you can use the `break` statement:

  ```go
    package main

    import "fmt"

    func main() {
        for i := 0; i < 5; i++ {
            if i == 3 {
            break
            }
            fmt.Println(i)
        }
    }
  ```

  - This program will print:

    ```bash
    0
    1
    2
    ```

- There is no `while` loop in Go, but you can mimic it using the `for` loop:

  ```go
    package main

    import "fmt"

    func main() {
        i := 0
        for i < 5 {
            fmt.Println(i)
            i++
        }
    }
  ```

  - This program will print:

    ```bash
    0
    1
    2
    3
    4
    ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `for_loops.go` which does the following:

1. Use a `for` loop to iterate 0 to 5.
2. In each iteration, check if the number is even. If the number is even, print the number. Otherwise, skip to the next iteration.

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
$ go run for_loops.go
```

The expected output of your program should be:

```bash
0
2
4
```

This challenge will help you practice using for loops, arrays, and conditional statements in Go. Good luck! 🍀
