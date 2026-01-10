# URL

## What is this topic?

The url package provides URL parsing and manipulation. This project covers parsing URLs and working with query parameters.

## Parse URL

```go
u, _ := url.Parse("https://example.com/path?key=value")
```

## Build URL

```go
u := &url.URL{
    Scheme: "https",
    Host:   "example.com",
    Path:   "/path",
}
```

## Query Parameters

```go
values := u.Query()
values.Set("key", "value")
u.RawQuery = values.Encode()
```

## Key Concepts

- **Parse**: Parse URL string
- **Query**: Access query parameters
- **Encode**: Encode query string
