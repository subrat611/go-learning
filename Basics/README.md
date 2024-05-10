# Basics - Packages, Functions, and Variables

## What is Go?

- It is statically typed, compiled programming language and syntactically similar to C.

## Packages

- Go uses packages (just like the header files in C++) to organize code in reusable modules, encapsulating functionality.
- All source files in a directory must share the same package name.

  ```mdx
  myproject/
  ├── main.go
  ├── mypkg/
  │ ├── math.go
  │ └── string.go
  ```

  ```go
  package main

  import (
      "fmt"
      "myproject/mypkg"
  )

  func main() {
      fmt.Println(utils.Add(3, 5))
  }

  ```

- The recommended style of naming in Go is that identifiers will be named using `camelCase`, except for those meant to be accessible across packages which should be `PascalCase`.

## Variables

- Go is statically-typed, which means all variables must have a defined type at compile-time.

- Variables can be defined by explicitly specifying a type

  ```go
  var explicit int // Explicitly typed
  ```

- If you want compiler to assign the variable type to match the type of the initializer. You can use initializer.

  ```go
  implicit := 10   // Implicitly typed as an int
  ```

## Constants

- The value can not change during the execution of the program.
- It defined using the `const` keyword.
  `const age = 21`

## Functions

- Go functions accept zero or more parameters.
- Parameters must be explicitly typed, **there is no type inference**.

  ```go
  package main

  func main() {
      fmt.Println(hi("user_name"))
  }

  func hi (name string) string {
    return "hi " + name
  }
  ```
