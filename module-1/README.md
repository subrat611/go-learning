# Concept 2: Variables, Constants & Data Types

- Go is a statically typed, compiled language.

This means:

- Every variable has a type at compile time
- Types cannot randomly change like in Python/JS
- Compiler catches type errors early (huge advantage)

## Variables in Go

### Using var keyword

Explicit declaration:

```go
var age int = 25
var name string = "Subrat"
var price float64 = 99.99
```

The type is written after the variable name.

### Type inference (Go figures out type automatically)

```go
var city = "Bangalore"   // type = string
var rating = 4.5         // type = float64
var active = true        // type = bool
```

The compiler infers the type from the value.

### Short variable declaration (:=) (type inference)

Most commonly used inside functions:

```go
count := 10
message := "Hello"
isLoggedIn := false
```

Rules of :=

- Only allowed inside functions
- Cannot redeclare existing variables:

```go
x := 10
x := 20   // ❌ error
x = 20    // ✅ assignment
```

## Constants in Go

Use const for values that must not change.

```go
const Pi = 3.14
const AppName = "ExpenseTracker"
```

**Rules of constants**

- Must be assigned at compile time
- Cannot be assigned using runtime values

```go
const x = math.Sqrt(4)  // ❌ invalid, because it's runtime
```

## Basic Data Types in Go

Numeric Types

| Type                    | Description                   |
| ----------------------- | ----------------------------- |
| `int`, `int32`, `int64` | whole numbers                 |
| `float32`, `float64`    | decimal numbers               |
| `uint`                  | unsigned (no negative values) |

```go
var name string = "GoLang"
var isActive bool = true
```

**Zero Values (super important)**

If you declare a variable without assigning:

```go
var age int
var name string
var price float64
var flag bool
```

| Type   | Default (Zero Value) |
| ------ | -------------------- |
| int    | `0`                  |
| float  | `0`                  |
| string | `""` (empty string)  |
| bool   | `false`              |

This avoids null pointer issues common in other languages.

# Concept 3: Conditionals (if, else, switch)

## if, else if, else

```go
if condition {
    // code
} else if anotherCondition {
    // code
} else {
    // code
}
```

## switch

```go
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Unknown")
}
```

## Inline variable declaration in if (Go-only feature)

- The variable `n` exists only inside the if block.

```go
if n := len(name); n > 5 {
    fmt.Println("Long name")
}
```

# Loops in Go (for)

- Go has only one loop keyword: `for`.

It can behave like:

- A traditional loop
- A while-loop
- An infinite loop

## Traditional C-style loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## While-style loop

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

## Infinite loop

```go
for {
    fmt.Println("Running forever...")
}
```

## Loop with break and continue

```go
for i := 1; i <= 10; i++ {
    if i == 5 {
        continue // skip printing 5
    }
    if i == 8 {
        break // stop the loop
    }
    fmt.Println(i)
}
```
