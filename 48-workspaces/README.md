# Workspaces

## What is this topic?

Go workspaces allow working with multiple modules. This project covers workspace setup and management.

## Workspace Setup

```bash
go work init
go work use ./module1
go work use ./module2
```

## go.work File

```
go 1.21

use (
    ./module1
    ./module2
)
```

## Key Concepts

- **go work init**: Initialize workspace
- **go work use**: Add modules
- **Multiple modules**: Work with several modules
