package main

import (
	"fmt"
	"os"
	"text/template"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Templates                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📄 TEMPLATES")
	fmt.Println("   tmpl := template.Must(template.New(\"test\").Parse(\"Hello, {{.}}!\"))")
	fmt.Println("   tmpl.Execute(os.Stdout, \"World\")")
	fmt.Println()
}

func exampleTemplate() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Template")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	tmpl := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	
	fmt.Println("Code:")
	fmt.Println("  tmpl := template.Must(template.New(\"test\").Parse(\"Hello, {{.}}!\"))")
	fmt.Println("  tmpl.Execute(os.Stdout, \"World\")")
	fmt.Println()
	fmt.Println("Output:")
	tmpl.Execute(os.Stdout, "World")
	fmt.Println()
	fmt.Println()
}

func main() {
	printTheory()
	exampleTemplate()
}
