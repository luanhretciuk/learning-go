# Unsafe

## What is this topic?

The unsafe package provides low-level operations. This project covers unsafe.Pointer and when to use unsafe.

## Unsafe Pointer

```go
import "unsafe"
var x int = 42
p := unsafe.Pointer(&x)
```

## Sizeof

```go
size := unsafe.Sizeof(x)
```

## Key Concepts

- **unsafe.Pointer**: Convert between pointer types
- **Use with caution**: Bypasses type safety
- **Rare use cases**: System programming, performance
