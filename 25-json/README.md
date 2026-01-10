# JSON

## What is this topic?

The encoding/json package provides JSON encoding and decoding. This project covers Marshal, Unmarshal, struct tags, and custom marshaling.

## Marshal

Convert Go value to JSON:

```go
data, err := json.Marshal(person)
```

## Unmarshal

Convert JSON to Go value:

```go
var person Person
err := json.Unmarshal(data, &person)
```

## Struct Tags

Control JSON field names:

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}
```

## Custom Marshaling

Implement `MarshalJSON` and `UnmarshalJSON`:

```go
func (p Point) MarshalJSON() ([]byte, error) { ... }
```

## Key Concepts

- **Marshal**: Go to JSON
- **Unmarshal**: JSON to Go
- **Struct tags**: Control field names
- **Custom marshaling**: Implement interfaces
