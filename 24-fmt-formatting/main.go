package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Fmt Formatting                           ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🖨️  PRINT FUNCTIONS")
	fmt.Println("   fmt.Print(\"hello\")")
	fmt.Println("   fmt.Println(\"hello\")")
	fmt.Println("   fmt.Printf(\"Value: %d\\n\", 42)")
	fmt.Println()
	fmt.Println("📝 FORMAT VERBS")
	fmt.Println("   %v - default format")
	fmt.Println("   %d - integer")
	fmt.Println("   %f - float")
	fmt.Println("   %s - string")
	fmt.Println("   %t - boolean")
	fmt.Println("   %p - pointer")
	fmt.Println("   %T - type")
	fmt.Println()
}

func exampleFormatVerbs() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Format Verbs")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	i := 42
	f := 3.14
	s := "hello"
	b := true
	
	fmt.Println("Code:")
	fmt.Println("  fmt.Printf(\"%d %f %s %t\", i, f, s, b)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  %d %f %s %t\n", i, f, s, b)
	fmt.Println()
}

func exampleSprintf() {
	fmt.Println("Example: Sprintf")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	s := fmt.Sprintf("Value: %d, Name: %s", 42, "Go")
	
	fmt.Println("Code:")
	fmt.Println("  s := fmt.Sprintf(\"Value: %d, Name: %s\", 42, \"Go\")")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  %s\n", s)
	fmt.Println()
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

func exampleCustomFormatting() {
	fmt.Println("Example: Custom Formatting (Stringer)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	p := Point{3, 4}
	
	fmt.Println("Code:")
	fmt.Println("  type Point struct { X, Y int }")
	fmt.Println("  func (p Point) String() string { ... }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  %v\n", p)
	fmt.Println()
}

func main() {
	printTheory()
	exampleFormatVerbs()
	exampleSprintf()
	exampleCustomFormatting()
}
