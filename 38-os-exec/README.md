# OS Exec

## What is this topic?

The os/exec package provides execution of external commands. This project covers running commands and capturing output.

## Run Command

```go
cmd := exec.Command("ls", "-l")
output, err := cmd.Output()
```

## Combined Output

```go
output, err := cmd.CombinedOutput()
```

## Key Concepts

- **Command**: Create command
- **Output**: Capture stdout
- **CombinedOutput**: Capture stdout and stderr
