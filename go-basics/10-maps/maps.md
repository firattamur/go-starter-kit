## 10 - Maps in Go:

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
