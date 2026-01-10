# HTTP Server

## What is this topic?

The net/http package provides HTTP client and server functionality. This project covers creating HTTP servers, handlers, and routing.

## Basic Server

```go
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
```

## Handler Function

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
```

## Custom Handler

```go
type MyHandler struct{}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { ... }
```

## Key Concepts

- **HandleFunc**: Register route handlers
- **ListenAndServe**: Start server
- **ResponseWriter**: Write response
- **Request**: Read request data
