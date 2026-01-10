# Interfaces

## What is this topic?

Interfaces define behavior through method signatures. This project covers interface definition, implicit implementation, empty interface, type assertions, and interface composition.

## Interface Definition

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

## Implicit Implementation

Types implement interfaces implicitly - no explicit declaration needed:

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
// Circle automatically implements Shape
```

## Empty Interface

The empty interface `interface{}` accepts any type:

```go
var i interface{} = 42
var s interface{} = "hello"
```

## Type Assertions

Extract concrete type from interface:

```go
var shape Shape = Circle{5}
circle := shape.(Circle)  // assertion
circle, ok := shape.(Circle)  // safe assertion
```

## Interface Composition

Interfaces can be composed:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

## Key Concepts

- **Implicit implementation**: No explicit declaration needed
- **Empty interface**: `interface{}` accepts any type
- **Type assertions**: Extract concrete types
- **Interface composition**: Build complex interfaces from simple ones
