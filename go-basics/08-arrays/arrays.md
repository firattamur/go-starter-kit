## 08 - Arrays in Go:

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
