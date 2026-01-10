# Testing

## What is this topic?

The testing package provides testing functionality. This project covers unit tests, table-driven tests, benchmarks, and examples.

## Basic Test

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

## Table-Driven Tests

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {2, 3, 5},
        {0, 0, 0},
    }
    for _, tt := range tests {
        result := Add(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Add(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
        }
    }
}
```

## Benchmarks

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

## Key Concepts

- **Test functions**: Must start with Test
- **Table-driven**: Test multiple cases
- **Benchmarks**: Measure performance
- **Examples**: Document usage
