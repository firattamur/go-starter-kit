## 08 - Switch Statements in Go:

- Go has a `switch` statement for conditional execution.

- Simple `switch` statement:

  ```go
    i := 2

    switch i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        default:
            fmt.Println("not one or two")
    }
  ```

- `switch` statements can also be used to initialize variables:

  ```go
    switch i := 2; i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        default:
            fmt.Println("not one or two")
    }
  ```

- There can be multiple expressions in a `case` statement:

  ```go
    switch i := 2; i {
        case 1, 2, 3:
            fmt.Println("one, two, or three")
        case 4, 5, 6:
            fmt.Println("four, five, or six")
        default:
            fmt.Println("another number")
    }
  ```

- We can also use the `fallthrough` keyword to execute the next `case` statement:

  ```go
    switch i := 2; i {
        case 1, 2, 3:
            fmt.Println("one, two, or three")
            fallthrough
        case 4, 5, 6:
            fmt.Println("four, five, or six")
        default:
            fmt.Println("another number")
    }
  ```

  - `fallthrough` is not the default behavior in Go. It must be explicitly stated.

  - A useful example of a `fallthrough` is when we want to check for a range of values:

  ```go
    switch i := 2; i {
        case 1, 2, 3:
            fmt.Println("one, two, or three")
            fallthrough
        case 4, 5, 6:
            fmt.Println("four, five, or six")
            fallthrough
        case 7, 8, 9:
            fmt.Println("seven, eight, or nine")
        default:
            fmt.Println("another number")
    }
  ```

  - Output:

    ```
    one, two, or three
    four, five, or six
    seven, eight, or nine
    ```

- `switch` statements can also be used to check the type of interface:

  ```go
      var i interface{} = 1

      switch i.(type) {
          case int:
              fmt.Println("i is an int")
          case float64:
              fmt.Println("i is a float64")
          case string:
              fmt.Println("i is a string")
          default:
              fmt.Println("i is another type")
      }
  ```

  - Output:

    ```
    i is an int
    ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `switch.go` which does the following:

1. Declare a `string` variable named `customerType` and assign a value to it (e.g., "regular", "premium", "vip", or any other string).

2. Write a `switch` statement that prints the following messages depending on the value of `customerType`:

   - "You have earned 10 points!" if `customerType` is "regular".

   - "You have earned 30 points!" if `customerType` is "premium".

   - "You have earned 50 points!" if `customerType` is "vip".

   - "You have earned 0 points!" if `customerType` is anything else.

You can start with the following code skeleton:

```go
package main

import "fmt"

func main() {

    customerType := "regular"
    // Your code here

}
```

After you have finished writing your code, you can run it with the following command:

```bash
$ go run switch.go
```

The expected output of your program should be:

```bash
You have earned 10 points!
```

This challenge will help you practice using switch statements in Go. Good luck! 🍀
