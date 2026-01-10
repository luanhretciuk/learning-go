# Functions

## What is this topic?

Functions are the building blocks of Go programs. This project covers function syntax, parameters, return values, variadic functions, and named returns.

## Function Syntax

### Basic Function

```go
func add(a int, b int) int {
    return a + b
}
```

### Function Components

- `func` - keyword to define a function
- `add` - function name
- `(a int, b int)` - parameters
- `int` - return type
- `{ ... }` - function body

## Multiple Return Values

Go functions can return multiple values:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## Named Returns

You can name return values:

```go
func calculate(x, y int) (sum int, product int) {
    sum = x + y
    product = x * y
    return  // naked return
}
```

## Variadic Functions

Functions that accept variable number of arguments:

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}
```

## Function Types

Functions are first-class values in Go:

```go
var operation func(int, int) int
operation = add
```

## Key Concepts

- **Function declaration**: `func name(params) returnType`
- **Multiple returns**: Go's unique feature
- **Named returns**: Can use naked return
- **Variadic functions**: `...` syntax
- **First-class functions**: Functions as values
