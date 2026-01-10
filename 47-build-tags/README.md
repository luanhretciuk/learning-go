# Build Tags

## What is this topic?

Build tags control which files are included in a build. This project covers build constraints and conditional compilation.

## Build Tags

```go
//go:build linux
// +build linux

package main
```

## Multiple Tags

```go
//go:build linux || darwin
// +build linux darwin
```

## Key Concepts

- **Build tags**: Control file inclusion
- **Platform-specific**: Different code per OS
- **Conditional compilation**: Include/exclude code
