# Race Detector

## What is this topic?

Go's race detector helps find data races. This project covers using the race detector and fixing race conditions.

## Using Race Detector

```bash
go run -race program.go
go test -race
```

## Race Condition Example

```go
var counter int
go func() { counter++ }()
go func() { counter++ }()
```

## Fixing Races

Use synchronization primitives:
- Mutex
- Channels
- Atomic operations

## Key Concepts

- **-race flag**: Enable race detector
- **Data races**: Concurrent access without synchronization
- **Fix**: Use proper synchronization
