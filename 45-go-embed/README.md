# Go Embed

## What is this topic?

The go:embed directive embeds files at compile time. This project covers embedding files and directories.

## Embed File

```go
//go:embed file.txt
var content string
```

## Embed Directory

```go
//go:embed static/*
var staticFiles embed.FS
```

## Key Concepts

- **go:embed**: Compile-time embedding
- **Files**: Embed single files
- **Directories**: Embed directory trees
- **embed.FS**: File system interface
