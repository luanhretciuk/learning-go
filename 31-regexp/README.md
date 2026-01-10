# Regexp

## What is this topic?

The regexp package provides regular expression functionality. This project covers compilation, matching, and submatch extraction.

## Compile

```go
re := regexp.MustCompile(`\d+`)
```

## Match

```go
matched := re.MatchString("123")
```

## Find

```go
match := re.FindString("abc123def")
```

## FindAll

```go
matches := re.FindAllString("a1b2c3", -1)
```

## Submatches

```go
re := regexp.MustCompile(`(\d+)-(\d+)`)
matches := re.FindStringSubmatch("123-456")
```

## Key Concepts

- **Compile**: Create regexp
- **Match**: Check if string matches
- **Find**: Extract matches
- **Submatches**: Extract groups
