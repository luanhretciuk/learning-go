# Advanced Patterns

## What is this topic?

Common Go patterns and idioms. This project covers builder pattern, functional options, worker pools, and pipelines.

## Builder Pattern

```go
type Builder struct {
    value string
}
func (b *Builder) SetValue(v string) *Builder {
    b.value = v
    return b
}
```

## Functional Options

```go
type Config struct {
    timeout time.Duration
}
func WithTimeout(d time.Duration) func(*Config) {
    return func(c *Config) { c.timeout = d }
}
```

## Worker Pool

```go
jobs := make(chan int, 100)
results := make(chan int, 100)
for w := 0; w < 3; w++ {
    go worker(jobs, results)
}
```

## Key Concepts

- **Builder**: Fluent API
- **Functional options**: Flexible configuration
- **Worker pools**: Concurrent processing
- **Pipelines**: Data processing chains
