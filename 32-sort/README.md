# Sort

## What is this topic?

The sort package provides sorting functionality. This project covers sorting slices and custom sorting.

## Basic Sorting

```go
import "sort"
ints := []int{3, 1, 4, 1, 5}
sort.Ints(ints)
```

## Custom Sorting

```go
type Person struct {
    Name string
    Age  int
}

sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

## Key Concepts

- **Ints/Strings**: Sort built-in types
- **Slice**: Sort with custom comparison
- **Stable**: Stable sort (preserves order of equal elements)
