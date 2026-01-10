# Methods

## What is this topic?

Methods are functions with a special receiver argument. This project covers value receivers vs pointer receivers, method sets, and method promotion.

## Method Syntax

```go
type Point struct {
    X, Y float64
}

func (p Point) Distance() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
```

## Value Receivers

Methods with value receivers operate on a copy:

```go
func (p Point) Move(dx, dy float64) {
    p.X += dx  // doesn't modify original
    p.Y += dy
}
```

## Pointer Receivers

Methods with pointer receivers can modify the original:

```go
func (p *Point) Move(dx, dy float64) {
    p.X += dx  // modifies original
    p.Y += dy
}
```

## Method Sets

- Value receiver: `(p Point)` - can be called on both value and pointer
- Pointer receiver: `(p *Point)` - can be called on both value and pointer
- Go automatically handles the conversion

## Key Concepts

- **Value receivers**: Operate on copy, cannot modify original
- **Pointer receivers**: Can modify original, more efficient for large structs
- **Method sets**: Automatic conversion between value and pointer
- **When to use**: Pointer receivers for methods that modify, value receivers for methods that don't
