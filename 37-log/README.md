# Log

## What is this topic?

The log package provides logging functionality. This project covers basic logging operations.

## Basic Logging

```go
log.Print("message")
log.Println("message")
log.Printf("format %s", "value")
```

## Log Levels

```go
log.SetPrefix("ERROR: ")
log.SetFlags(log.Ldate | log.Ltime)
```

## Key Concepts

- **Print/Println/Printf**: Log messages
- **SetPrefix**: Add prefix to logs
- **SetFlags**: Configure log format
