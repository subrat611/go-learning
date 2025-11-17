# Concept 1: Functions in Go

- Functions are the heart of clean Go code.
- Everything in Go is organized around small, reusable functions.

## What is a Function?

A function in Go:

- Takes input (Optional)
- Does some work
- Returns output (Optional)

**Syntax**

```go
func functionName(param1 type, param2 type) returnType {
    return value
}
```

**Example**

```go
func add(a int, b int) int {
    return a + b
}
```

## Return multiple values (Go is famous for this)

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("Division by zero")
    }

    return a/b, nil
}
```

## Named returns (Optinal feature)

```go
func rectangle(w, h float64) (area float64) {
    area = w * h
    return // return area implicity
}
```

## Variadic functions

- A function that accepts any number of arguments

```go
func sum(nums ...int) int {
    total := 0

    for _, v := range nums {
        total += v
    }

    return total
}
```

## Functions are first-class

- You can store functions in variables

```go
f := func() {
    fmt.Println("Hello")
}

f()
```

# Concept 2: Arrays & Slices

Arrays in Go:

- Have fixed size
- Size is part of the type
- Cannot grow or shrink
- Usually avoided in real Go code

**Example**

```go
var num [3]int
num[0] = 10
num[1] = 20
num[2] = 30

// or

nums := [3]int{10, 20, 30}
```

Arrays are mostly used internally, not in application level code.

## Slices (Used everywhere)

Slices are:

- Dynamic
- Flexible
- Backed by arrays
- The most used data structure in Go

**Syntax**

```go
nums := []int{1, 2, 3, 4}

// or empty slice
var scores []int // nil slice

// or make based slice
scores := make([]int, 3) // [0 0 0] with length 3 and capacity 3
```

### Slice Properties: Length & Capacity

- Capacity expands when you use `append`.

```go
nums := []int{1, 2, 3}
fmt.Println(len(nums))
fmt.Println(cap(nums))
```

### Slicing Syntax

```go
nums := []int{10, 20, 30, 40, 50}

fmt.Println(nums[1:4]) // [20 30 40]
fmt.Println(nums[:3]) // [10 20 30]
fmt.Println(nums[2:]) // [20 40 50]
fmt.Println(nums[:]) // whole slice
```

### Append (the most important slice operation)

- Appending increases length and may change capacity

```go
nums := []int{1, 2, 3}
nums = append(nums, 4)
```

### Delete from slice (Common pattern)

- Go doesn't have a delete keyword for slices. Instead, you rebuild the slice without the unwanted element.

```go
// delete index i
nums = append(nums[:i], nums[i+1:]...)
```

### Looping through slices with range

```go
nums := []int{10, 20, 30}

for index, value := range nums {
    fmt.Println(index, value)
}

// if you want only values
for _, value := range nums {
    fmt.Println(value)
}
```
