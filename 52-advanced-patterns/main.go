package main

import "fmt"

type Builder struct {
	value string
}

func (b *Builder) SetValue(v string) *Builder {
	b.value = v
	return b
}

func (b *Builder) Build() string {
	return b.value
}

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Advanced Patterns                        ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🎯 COMMON PATTERNS")
	fmt.Println("   • Builder pattern")
	fmt.Println("   • Functional options")
	fmt.Println("   • Worker pools")
	fmt.Println("   • Pipelines")
	fmt.Println()
}

func exampleBuilder() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Builder Pattern")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	result := new(Builder).SetValue("Hello").SetValue("World").Build()
	
	fmt.Println("Code:")
	fmt.Println("  new(Builder).SetValue(\"Hello\").SetValue(\"World\").Build()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Result: %s\n", result)
	fmt.Println()
}

func main() {
	printTheory()
	exampleBuilder()
}
