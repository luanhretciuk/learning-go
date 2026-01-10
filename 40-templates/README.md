# Templates

## What is this topic?

The text/template and html/template packages provide template processing. This project covers template syntax and execution.

## Basic Template

```go
tmpl := template.Must(template.New("test").Parse("Hello, {{.}}!"))
tmpl.Execute(os.Stdout, "World")
```

## Template Variables

```go
tmpl := template.Must(template.New("test").Parse("Name: {{.Name}}"))
tmpl.Execute(os.Stdout, Person{Name: "Alice"})
```

## Key Concepts

- **Parse**: Parse template string
- **Execute**: Execute template with data
- **Variables**: Access data with {{.}}
- **html/template**: HTML-safe templates
