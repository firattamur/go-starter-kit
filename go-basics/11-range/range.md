## 11 - Range in Go:

- The `range` keyword can be used to iterate over an array, slice, string, map, or channel.

- A simple example of using `range` to iterate over an array:

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

- A simple example of using `range` to iterate over a slice:

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

- A simple example of using `range` to iterate over a string:

  ```go
      var s = "Hello, World!"

      for i, v := range s {
          fmt.Println(i, v)
      }
  ```

  - Output:

    ```bash
    0 72
    1 101
    2 108
    3 108
    4 111
    5 44
    6 32
    7 87
    8 111
    9 114
    10 108
    11 100
    12 33
    ```

    - The second value returned by `range` is the `rune` value of the character.

    - A `rune` is an alias for `int32` and is used to represent a Unicode code point.

    - To get the character value, we can use the `string` function:

      ```go
          var s = "Hello, World!"

          for i, v := range s {
              fmt.Println(i, string(v))
          }
      ```

      - Output:

        ```bash
        0 H
        1 e
        2 l
        3 l
        4 o
        5 ,
        6
        7 W
        8 o
        9 r
        10 l
        11 d
        12 !
        ```

- A simple example of using `range` to iterate over a map:

  ```go
      var m = map[string]int{"foo": 42, "bar": 27}

      for k, v := range m {
          fmt.Println(k, v)
      }
  ```

  - Output:

    ```bash
    foo 42
    bar 27
    ```

- A simple example of using `range` to iterate keys of a map:

  ```go
      var m = map[string]int{"foo": 42, "bar": 27}

      for k := range m {
          fmt.Println(k)
      }
  ```

  - Output:

    ```bash
    foo
    bar
    ```

## 🚀 Coding Challenge:

Your task is to write a Go program named `range.go` which does the following:

1. Declare a `string` variable named `sentence` and assign the value "Go is awesome!" to it.

2. Use the `range` keyword to iterate over the `sentence` string and print each character.

3. Declare a `map` named `m` that maps `string` keys to `int` values.

4. Map sentence characters to their respective ASCII values to the `m` map.

5. Use the `range` keyword to iterate over the `m` map and print each key, value pair.

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
$ go run range.go
```

The expected output of your program should be:

```bash
0 G
1 o
2
3 i
4 s
5
6 a
7 w
8 e
9 s
10 o
11 m
12 e
13 !
m 11
! 13
  5
i 3
a 6
e 12
G 0
o 10
s 9
w 7
```

This challenge will help you practice using the `range` keyword in Go. Good luck! 🍀
