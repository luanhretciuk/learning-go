# Net TCP/UDP

## What is this topic?

The net package provides low-level networking. This project covers TCP and UDP connections.

## TCP Server

```go
ln, _ := net.Listen("tcp", ":8080")
conn, _ := ln.Accept()
```

## TCP Client

```go
conn, _ := net.Dial("tcp", "localhost:8080")
```

## Key Concepts

- **TCP**: Reliable, connection-oriented
- **UDP**: Fast, connectionless
- **Listen**: Create server
- **Dial**: Create client connection
