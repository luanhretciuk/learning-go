# Arrays & Slices

## What is this topic?

Arrays and slices are fundamental data structures in Go. This project covers the differences between arrays and slices, slice operations, and how to work with them effectively.

## Arrays

Arrays are fixed-size sequences:

```go
var arr [5]int           // array of 5 integers
arr := [5]int{1, 2, 3}    // initialized array
arr := [...]int{1, 2, 3}  // size inferred
```

### Key Points

- Fixed size (cannot grow)
- Value type (copied when assigned)
- Size is part of the type

## Slices

Slices are dynamic arrays:

```go
var s []int              // nil slice
s := []int{1, 2, 3}      // slice literal
s := make([]int, 5)      // make with length
s := make([]int, 5, 10)  // make with length and capacity
```

### Slice Internals

A slice is a reference to an underlying array:
- **Pointer**: Points to underlying array
- **Length**: Number of elements
- **Capacity**: Maximum elements without reallocation

## Slice Operations

### Append

```go
s := []int{1, 2, 3}
s = append(s, 4, 5)  // [1, 2, 3, 4, 5]
```

### Slicing

```go
s := []int{0, 1, 2, 3, 4}
s[1:3]  // [1, 2] - from index 1 to 3 (exclusive)
s[:3]   // [0, 1, 2] - from start to 3
s[2:]   // [2, 3, 4] - from 2 to end
```

### Length and Capacity

```go
len(s)  // current length
cap(s)  // capacity
```

## Key Concepts

- **Arrays**: Fixed size, value type
- **Slices**: Dynamic, reference type
- **Append**: Add elements (may reallocate)
- **Slicing**: Create new slice from existing
- **Length vs Capacity**: Current size vs allocated size
