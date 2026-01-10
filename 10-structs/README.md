# Structs

## What is this topic?

Structs are composite data types that group together related fields. This project covers struct definition, initialization, field access, and embedded structs.

## Struct Definition

```go
type Person struct {
    Name string
    Age  int
}
```

## Struct Initialization

```go
p1 := Person{"Alice", 30}              // positional
p2 := Person{Name: "Bob", Age: 25}      // named fields
p3 := Person{Name: "Charlie"}           // partial (Age = 0)
var p4 Person                            // zero value
```

## Field Access

```go
p.Name = "David"
age := p.Age
```

## Embedded Structs

Structs can embed other structs:

```go
type Employee struct {
    Person
    Salary float64
}
```

## Key Concepts

- **Struct definition**: `type Name struct { ... }`
- **Initialization**: Positional or named fields
- **Field access**: Dot notation
- **Embedded structs**: Composition over inheritance
