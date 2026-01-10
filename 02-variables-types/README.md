# Variables & Types

## What is this topic?

This project covers how to declare and use variables in Go, the type system, constants, and zero values. Understanding variables is fundamental to programming in Go.

## Variable Declaration

### Three Ways to Declare Variables

1. **Using `var` keyword:**
   ```go
   var name string = "Go"
   ```

2. **Short variable declaration (most common):**
   ```go
   name := "Go"
   ```

3. **Multiple variables:**
   ```go
   var x, y int = 1, 2
   a, b := 3, 4
   ```

### Type Inference

Go can infer the type from the value:

```go
var name = "Go"  // type is string
age := 25        // type is int
```

## Basic Types

### Numeric Types

- **Integers:**
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`
  - `byte` (alias for uint8)
  - `rune` (alias for int32, represents Unicode code point)

- **Floating Point:**
  - `float32`, `float64`

- **Complex:**
  - `complex64`, `complex128`

### Other Types

- `bool` - true or false
- `string` - sequence of bytes (UTF-8 encoded)
- `error` - error type

## Zero Values

In Go, variables are automatically initialized to their "zero value" if not explicitly set:

- Numbers: `0`
- Booleans: `false`
- Strings: `""` (empty string)
- Pointers, slices, maps, channels, functions, interfaces: `nil`

## Constants

Constants are declared with `const`:

```go
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

### iota

`iota` is a special constant generator:

```go
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
)
```

## Type Conversion

Go requires explicit type conversion:

```go
var i int = 42
var f float64 = float64(i)
```

## Key Concepts

- **Variable declaration**: `var`, `:=`, or multiple
- **Type system**: Static typing with type inference
- **Zero values**: Automatic initialization
- **Constants**: Immutable values
- **iota**: Constant generator
- **Type conversion**: Explicit conversion required
