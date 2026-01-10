package main

import (
	"fmt"
	"strings"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Strings & Runes                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📝 STRINGS")
	fmt.Println("   Strings in Go are UTF-8 encoded byte sequences")
	fmt.Println("   • Immutable (cannot change after creation)")
	fmt.Println("   • len(string) returns byte length, not character count")
	fmt.Println()
	fmt.Println("🔤 RUNES")
	fmt.Println("   A rune is an alias for int32")
	fmt.Println("   • Represents a Unicode code point")
	fmt.Println("   • Use []rune for character-level operations")
	fmt.Println()
	fmt.Println("🌐 UTF-8 ENCODING")
	fmt.Println("   Go strings are UTF-8 by default")
	fmt.Println("   • Each character can be 1-4 bytes")
	fmt.Println("   • Supports all Unicode characters")
	fmt.Println()
	fmt.Println("📚 STRINGS PACKAGE")
	fmt.Println("   Rich string manipulation functions:")
	fmt.Println("   • Contains, HasPrefix, HasSuffix")
	fmt.Println("   • Index, Replace, Split, Join")
	fmt.Println("   • ToUpper, ToLower, TrimSpace")
	fmt.Println()
}

func exampleStrings() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Strings")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	s1 := "Hello"
	s2 := "世界"
	s3 := s1 + " " + s2
	
	fmt.Println("Code:")
	fmt.Println("  s1 := \"Hello\"")
	fmt.Println("  s2 := \"世界\"")
	fmt.Println("  s3 := s1 + \" \" + s2")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  s1: %s, len: %d bytes\n", s1, len(s1))
	fmt.Printf("  s2: %s, len: %d bytes\n", s2, len(s2))
	fmt.Printf("  s3: %s, len: %d bytes\n", s3, len(s3))
	fmt.Println()
}

func exampleRunes() {
	fmt.Println("Example: Runes")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	s := "Hello, 世界"
	runes := []rune(s)
	
	fmt.Println("Code:")
	fmt.Println("  s := \"Hello, 世界\"")
	fmt.Println("  runes := []rune(s)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  String: %s\n", s)
	fmt.Printf("  Byte length: %d\n", len(s))
	fmt.Printf("  Rune count: %d\n", len(runes))
	fmt.Printf("  Runes: %v\n", runes)
	fmt.Println()
}

func exampleStringsPackage() {
	fmt.Println("Example: strings Package")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	s := "  Hello, World!  "
	
	fmt.Println("Code:")
	fmt.Println("  s := \"  Hello, World!  \"")
	fmt.Println("  strings.Contains(s, \"World\")")
	fmt.Println("  strings.HasPrefix(s, \"  Hello\")")
	fmt.Println("  strings.TrimSpace(s)")
	fmt.Println("  strings.ToUpper(s)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Contains \"World\": %t\n", strings.Contains(s, "World"))
	fmt.Printf("  HasPrefix \"  Hello\": %t\n", strings.HasPrefix(s, "  Hello"))
	fmt.Printf("  TrimSpace: \"%s\"\n", strings.TrimSpace(s))
	fmt.Printf("  ToUpper: \"%s\"\n", strings.ToUpper(s))
	fmt.Println()
}

func exampleStringBuilder() {
	fmt.Println("Example: strings.Builder")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	result := builder.String()
	
	fmt.Println("Code:")
	fmt.Println("  var builder strings.Builder")
	fmt.Println("  builder.WriteString(\"Hello\")")
	fmt.Println("  builder.WriteString(\" \")")
	fmt.Println("  builder.WriteString(\"World\")")
	fmt.Println("  result := builder.String()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Result: \"%s\"\n", result)
	fmt.Println()
	fmt.Println("💡 Tip: Use strings.Builder for efficient string concatenation")
	fmt.Println()
}

func main() {
	printTheory()
	exampleStrings()
	exampleRunes()
	exampleStringsPackage()
	exampleStringBuilder()
}
