![](/.github/assets/logo.png)

A starter kit for those interested in learning Go. This repository contains a collection of resources, mini-projects, and code examples to help you get up to speed with Go.

## Table of Contents

1. [Introduction](#introduction)
2. [Getting Started](#getting-started)
3. [Go Basics](#go-basics)
4. [Go Intermediate](#intermediate-go)
5. [Go Advanced](#advanced-go)
6. [Tools and Libraries](#go-tools-and-libraries)
7. [Mini-Projects](#mini-projects)
8. [Contributing](#contributing)
9. [Resources](#resources)

## 1. Introduction:

Learn about the history of Go, its benefits, and why you might want to use it.

## 2. Getting Started:

Get Go installed on your machine and set up your development environment.

## 3. Go Basics:

Dive into Go syntax with hands-on examples. Each subdirectory in this section contains a README with explanations and code examples for a particular topic, such as variables, functions, or control structures.

<details>
    <summary>01 - Hello World in Go:</summary>

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

### 🚀 Coding Challenge:

Your task is to create a Go program that prints "Hello, Go!" instead of "Hello, World!". You should save this program as `hello_go.go` and run it using the `go run` command. Here is a skeleton of the code to get you started:

```go
package main

import "fmt"

func main() {
// Your code here
}
```

After you have finished writing your code, you can run it with the following command:

```bash
$ go run hello_go.go
```

Your program should output:

```bash
Hello, Go!
```

The goal here is to practice using the `fmt` package to print messages to the console. Good luck!

</details>

<details>
    <summary>02 - Data Types in Go:</summary>

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

</details>

<details>
    <summary>03 - Variables in Go:</summary>

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

</details>

<details>
    <summary>04 - Constants in Go:</summary>

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

</details>

<details>
    <summary>05 - For Loops in Go:</summary>

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

</details>

<details>
    <summary>06 - If/Else Statements in Go:</summary>

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

</details>

<details>
    <summary>07 - Switch Statements in Go:</summary>

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

</details>

<details>
    <summary>08 - Arrays in Go:</summary>

- Arrays are fixed-length sequences of elements of a single type.

- Arrays are declared using brackets:

  ```go
    var a [5]int
  ```

- Arrays are zero-indexed:

  ```go
    a[0] = 1
    a[1] = 2
    a[2] = 3
    a[3] = 4
    a[4] = 5
  ```

- Arrays can also be initialized with values:

  ```go
    var a = [5]int{1, 2, 3, 4, 5}
  ```

- Arrays can also be initialized without specifying the length:

  ```go
    var a = [...]int{1, 2, 3, 4, 5}
  ```

- Arrays can also be initialized using indexes:

  ```go
    var a = [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
  ```

- Arrays initialized with default value of 0, if not specified:

  ```go
    var a = [5]int{1, 2, 3}
  ```

  - Output:

    ```bash
    [1 2 3 0 0]
    ```

- Arrays can also be initialized using the `new` keyword:

  ```go
      var a = new([5]int)
  ```

  - Output:

    ```bash
    &[0 0 0 0 0]
    ```

    - `new` returns a pointer to the array.

    - We can access the elements of the array using the pointer:

      ```go
        var a = new([5]int)
        a[0] = 1
        a[1] = 2
        a[2] = 3
        a[3] = 4
        a[4] = 5

        fmt.Println(a)
      ```

      - Output:

        ```bash
        &[1 2 3 4 5]
        ```

- Arrays can also be declared using the `:=` operator:

  ```go
    a := [5]int{1, 2, 3, 4, 5}
  ```

- To get the length of an array, we can use the `len` function:

  ```go
    var a = [5]int{1, 2, 3, 4, 5}
    fmt.Println(len(a))
  ```

  - Output:

    ```bash
    5
    ```

- We can also declare multi-dimensional arrays:

  ```go
    var a = [2][2]int{{1, 2}, {3, 4}}
  ```

  - Output:

    ```bash
    [[1 2] [3 4]]
    ```

- To access the elements of a multi-dimensional array, we can use the following syntax:

  ```go
    var a = [2][2]int{{1, 2}, {3, 4}}
    fmt.Println(a[0][0])
    fmt.Println(a[0][1])
    fmt.Println(a[1][0])
    fmt.Println(a[1][1])
  ```

  - Output:

    ```bash
    1
    2
    3
    4
    ```

- We can also use the `for` loop to iterate over an array:

  ```go
      var a = [5]int{1, 2, 3, 4, 5}

      for i := 0; i < len(a); i++ {
          fmt.Println(a[i])
      }
  ```

  - Output:

    ```bash
    1
    2
    3
    4
    5
    ```

- We can also use the `range` keyword to iterate over an array:

  ```go
      var a = [5]int{1, 2, 3, 4, 5}

      for i, v := range a {
          fmt.Println(i, v)
      }
  ```

  - Output:

    ```bash
    0 1
    1 2
    2 3
    3 4
    4 5
    ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `arrays.go` which does the following:

1. Declare an array of 5 `int` elements named `a` and assign the values 1, 2, 3, 4, 5 to it.

2. Print the array.

3. Print the length of the array.

4. Declare a multi-dimensional array of 2x2 `int` elements named `b` and assign the values 1, 2, 3, 4 to it.

5. Print the multi-dimensional array.

6. Print the length of the multi-dimensional array.

7. Use a `for` loop to iterate over the array and print each element.

8. Use a `for` loop to iterate over the multi-dimensional array and print each element.

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
$ go run arrays.go
```

The expected output of your program should be:

```bash
[1 2 3 4 5]
5
[[1 2] [3 4]]
2
1
2
3
4
5
1
2
3
4
```

This challenge will help you practice using arrays in Go. Good luck! 🍀

</details>

<details>
    <summary>09 - Slices in Go:</summary>

- Slices are a built-in data structure in Go.

- Slices are similar to arrays, but they are dynamic.

- Slices are zero-indexed.

- Slices can be initialized using the `make` keyword:

  ```go
    var a = make([]int, 5)
  ```

- To get the length of a slice, we can use the built-in `len` function:

  ```go
    var a = make([]int, 5)
    fmt.Println(len(a))
  ```

  - Output:

    ```bash
    5
    ```

- To get the capacity of a slice, we can use the built-in `cap` function:

  ```go
      var a = make([]int, 5)
      fmt.Println(cap(a))
  ```

  - Output:

    ```bash
    5
    ```

- We can also declare multi-dimensional slices:

  ```go
      var a = make([][]int, 5)
  ```

- We can also declare slices using the `:=` operator:

  ```go
    a := []int{1, 2, 3, 4, 5}
  ```

- Additional functions for slices:

  - `append`:

    ```go
      var a = []int{1, 2, 3, 4, 5}
      a = append(a, 6)
      fmt.Println(a)
    ```

    - Output:

      ```bash
      [1 2 3 4 5 6]
      ```

  - `copy`:

    ```go
      var a = []int{1, 2, 3, 4, 5}
      var b = make([]int, 5)
      copy(b, a)
      fmt.Println(b)
    ```

    - Output:

      ```bash
      [1 2 3 4 5]
      ```

  - `range`:

    ```go
      var a = []int{1, 2, 3, 4, 5}
      for i, v := range a {
          fmt.Println(i, v)
      }
    ```

    - Output:

      ```bash
      0 1
      1 2
      2 3
      3 4
      4 5
      ```

- To get a slice of a slice, we can use the following syntax:

  ```go
    var a = []int{1, 2, 3, 4, 5}
    fmt.Println(a[1:3])
  ```

  - Output:

    ```bash
    [2 3]
    ```

- The cap and length of a slice can be different:

  ```go
    var a = make([]int, 5, 10)
    fmt.Println(len(a))
    fmt.Println(cap(a))
  ```

  - Output:

    ```bash
    5
    10
    ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `slices.go` which does the following:

1. Declare a slice of 5 `int` elements named `a` and assign the values 1, 2, 3, 4, 5 to it.

2. Print the slice.

3. Print the length of the slice.

4. Print the capacity of the slice.

5. Use the `append` function to add the value 6 to the slice.

6. Print the slice.

7. Use the `copy` function to copy the slice to a new slice named `c`.

8. Print the new slice.

9. Use a `for` loop to iterate over the slice and print each element.

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
$ go run slices.go
```

The expected output of your program should be:

```bash
[1 2 3 4 5]
5
5
[1 2 3 4 5 6]
[1 2 3 4 5]
0 1
1 2
2 3
3 4
4 5
5 6
```

This challenge will help you practice using slices in Go. Good luck! 🍀

</details>

<details>
    <summary>10 - Maps in Go:</summary>

- Maps are unordered collections of key-value pairs.

- Maps are declared using brackets:

  ```go
    var m map[string]int
  ```

- Maps are initialized using the `make` keyword:

  ```go
    var m = make(map[string]int)
  ```

- Maps can also be initialized using the `:=` operator:

  ```go
    m := map[string]int{"foo": 42}
  ```

- Maps can also be initialized using the `new` keyword:

  ```go
      var m = new(map[string]int)
  ```

- Maps can be accessed using brackets:

  ```go
    m := map[string]int{"foo": 42}
    fmt.Println(m["foo"])
  ```

  - Output:

    ```bash
    42
    ```

- Maps can be updated using brackets:

  ```go
    m := map[string]int{"foo": 42}
    m["foo"] = 27
    fmt.Println(m["foo"])
  ```

  - Output:

    ```bash
    27
    ```

- Maps can be deleted using the `delete` keyword:

  ```go
      m := map[string]int{"foo": 42}
      delete(m, "foo")
      fmt.Println(m["foo"])
  ```

  - Output:

    ```bash
    0
    ```

- Maps can be iterated using the `range` keyword:

  ```go
    m := map[string]int{"foo": 42, "bar": 27}
    for k, v := range m {
        fmt.Println(k, v)
    }
  ```

  - Output:

    ```bash
    foo 42
    bar 27
    ```

- Maps' keys can be iterated using the `range` keyword:

  ```go
    m := map[string]int{"foo": 42, "bar": 27}
    for k := range m {
        fmt.Println(k)
    }
  ```

  - Output:

    ```bash
    foo
    bar
    ```

- The length of a map can be obtained using the `len` function:

  ```go
    m := map[string]int{"foo": 42, "bar": 27}
    fmt.Println(len(m))
  ```

  - Output:

    ```bash
    2
    ```

- If a key does not exist in a map, the value returned is the zero value of the type:

  ```go
    m := map[string]int{"foo": 42, "bar": 27}
    fmt.Println(m["baz"])
  ```

  - Output:

    ```bash
    0
    ```

- To check if a key exists in a map, we can use the `ok` idiom:

  ```go
    m := map[string]int{"foo": 42, "bar": 27}
    if v, ok := m["foo"]; ok {
        fmt.Println(v)
    }
  ```

  - Output:

    ```bash
    42
    ```

- Maps are passed by reference:

  ```go
      m := map[string]int{"foo": 42, "bar": 27}
      fmt.Println(m)
      updateMap(m)
      fmt.Println(m)

    func updateMap(m map[string]int) {
        m["foo"] = 27
    }
  ```

  - Output:

    ```bash
    map[bar:27 foo:42]
    map[bar:27 foo:27]
    ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `maps.go` which does the following:

1. Declare a map named `m` that maps `string` keys to `int` values.

2. Initialize `m` with the following key-value pairs:

   - "foo" -> 42
   - "bar" -> 27

3. Print the map.

4. Print the length of the map.

5. Print key, value pairs of the map using the `range` keyword.

6. Delete the key "foo" from the map.

7. Print the map.

8. Print the length of the map.

9. Print the value of the key "foo" from the map.

10. Print `bar, 27` if the key "bar" exists in the map.

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
$ go run maps.go
```

The expected output of your program should be:

```bash
map[bar:27 foo:42]
2
bar 27
foo 42
map[bar:27]
1
0
bar, 27
```

This challenge will help you practice using maps in Go. Good luck! 🍀

</details>

## 4. Go Intermediate:

Explore more complex Go topics like interfaces, goroutines, and error handling. Each topic has its own subdirectory with a README and code examples.

## 5. Go Advanced:

Delve into advanced Go concepts and features. This section covers topics like reflection, cgo, and the unsafe package, among others.

## 6. Tools and Libraries

Discover the Go toolchain and various libraries that make development in Go a breeze. Learn about tools for testing, linting, formatting, and more.

## 7. Mini-Projects

Apply what you've learned in a series of mini-projects. These projects integrate multiple concepts and give you experience with end-to-end development in Go.

## 8. Resources

Find additional resources for learning Go, such as books, tutorials, blogs, and official Go documentation.

- [Go by Example](https://gobyexample.com/)

## 9. Contributing

Contributions are welcome! Create a pull request to add a new resource, fix a typo, or add a new mini-project.

## 10. Happy Coding!
