# OS Signal

## What is this topic?

The os/signal package provides signal handling. This project covers catching and handling OS signals.

## Signal Handling

```go
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt, syscall.SIGTERM)
sig := <-c
```

## Key Concepts

- **Notify**: Register signal channel
- **Interrupt**: Catch interrupt signals
- **Graceful shutdown**: Handle termination signals
