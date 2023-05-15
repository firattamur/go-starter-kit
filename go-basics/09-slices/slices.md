## 09 - Slices in Go:

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
