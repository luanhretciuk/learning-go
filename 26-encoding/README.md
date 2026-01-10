# Encoding

## What is this topic?

Go provides various encoding packages. This project covers gob (binary), xml, base64, and hex encoding.

## Gob (Binary)

```go
import "encoding/gob"
enc := gob.NewEncoder(writer)
dec := gob.NewDecoder(reader)
```

## XML

```go
import "encoding/xml"
data, err := xml.Marshal(value)
err := xml.Unmarshal(data, &value)
```

## Base64

```go
import "encoding/base64"
encoded := base64.StdEncoding.EncodeToString(data)
decoded, err := base64.StdEncoding.DecodeString(encoded)
```

## Hex

```go
import "encoding/hex"
encoded := hex.EncodeToString(data)
decoded, err := hex.DecodeString(encoded)
```

## Key Concepts

- **Gob**: Go-specific binary encoding
- **XML**: XML encoding/decoding
- **Base64**: Base64 encoding
- **Hex**: Hexadecimal encoding
