# Time Package

## What is this topic?

The time package provides time and date functionality. This project covers Time, Duration, formatting, parsing, and timers.

## Time

```go
now := time.Now()
t := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
```

## Duration

```go
d := 5 * time.Second
d := 10 * time.Minute
```

## Formatting

```go
t.Format("2006-01-02 15:04:05")
```

## Parsing

```go
t, _ := time.Parse("2006-01-02", "2023-01-01")
```

## Timers and Tickers

```go
timer := time.NewTimer(1 * time.Second)
ticker := time.NewTicker(1 * time.Second)
```

## Key Concepts

- **Time**: Represents a point in time
- **Duration**: Time intervals
- **Formatting**: Custom date/time formats
- **Timers/Tickers**: Time-based events
