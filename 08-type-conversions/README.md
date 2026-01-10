# Type Conversions

## What is this topic?

Go requires explicit type conversions. This project covers type conversions, type assertions, and type switches for working with interfaces.

## Type Conversions

Go requires explicit conversion between types:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(i)
```

## Type Assertions

Used with interfaces to extract concrete types:

```go
var i interface{} = "hello"
s := i.(string)        // assertion
s, ok := i.(string)    // safe assertion
```

## Type Switches

Switch on the type of an interface value:

```go
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown")
}
```

## Key Concepts

- **Type conversion**: Explicit conversion required
- **Type assertion**: Extract concrete type from interface
- **Type switch**: Switch on interface type
- **Safe assertions**: Check if assertion succeeds
