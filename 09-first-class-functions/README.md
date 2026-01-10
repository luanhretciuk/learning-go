# First-Class Functions

## What is this topic?

Functions are first-class values in Go. This project covers functions as values, function types, closures, and method expressions.

## Functions as Values

Functions can be assigned to variables:

```go
add := func(a, b int) int {
    return a + b
}
result := add(3, 4)
```

## Function Types

You can define function types:

```go
type Operation func(int, int) int

var op Operation = func(a, b int) int {
    return a + b
}
```

## Closures

Functions can capture variables from their enclosing scope:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Method Expressions

Methods can be called as functions:

```go
type Point struct{ X, Y int }
func (p Point) Distance() float64 { ... }

dist := Point.Distance  // method expression
result := dist(Point{1, 2})
```

## Key Concepts

- **First-class functions**: Functions as values
- **Function types**: Type definitions for functions
- **Closures**: Functions that capture variables
- **Method expressions**: Methods as functions
