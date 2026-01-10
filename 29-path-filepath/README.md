# Path & Filepath

## What is this topic?

The path and path/filepath packages provide path manipulation. This project covers path operations and filepath utilities.

## Path Package

```go
import "path"
base := path.Base("/usr/bin/go")
dir := path.Dir("/usr/bin/go")
ext := path.Ext("file.txt")
```

## Filepath Package

```go
import "path/filepath"
abs, _ := filepath.Abs("file.txt")
join := filepath.Join("dir", "file.txt")
```

## Key Concepts

- **path**: URL/path manipulation
- **filepath**: OS-specific path handling
- **Join**: Build paths safely
- **Base/Dir**: Extract path components
