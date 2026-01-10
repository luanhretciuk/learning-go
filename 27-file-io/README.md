# File I/O

## What is this topic?

File operations in Go. This project covers reading and writing files using os, io, and bufio packages.

## Reading Files

```go
data, err := os.ReadFile("file.txt")
```

## Writing Files

```go
err := os.WriteFile("file.txt", data, 0644)
```

## Buffered I/O

```go
file, _ := os.Open("file.txt")
reader := bufio.NewReader(file)
line, _ := reader.ReadString('\n')
```

## Key Concepts

- **os.ReadFile/WriteFile**: Simple file operations
- **bufio**: Buffered I/O for efficiency
- **File permissions**: Unix-style permissions (0644)
