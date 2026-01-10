package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Bytes & Strconv                           ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📦 BYTES PACKAGE")
	fmt.Println("   bytes.Buffer for efficient string building")
	fmt.Println()
	fmt.Println("🔢 STRCONV PACKAGE")
	fmt.Println("   String conversions")
	fmt.Println("   strconv.Itoa(42)")
	fmt.Println("   strconv.Atoi(\"42\")")
	fmt.Println()
}

func exampleBytesBuffer() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Bytes Buffer")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteString(" ")
	buf.WriteString("World")
	result := buf.String()
	
	fmt.Println("Code:")
	fmt.Println("  var buf bytes.Buffer")
	fmt.Println("  buf.WriteString(\"Hello\")")
	fmt.Println("  result := buf.String()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Result: %s\n", result)
	fmt.Println()
}

func exampleStrconv() {
	fmt.Println("Example: Strconv")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	s := strconv.Itoa(42)
	i, _ := strconv.Atoi("42")
	f, _ := strconv.ParseFloat("3.14", 64)
	
	fmt.Println("Code:")
	fmt.Println("  s := strconv.Itoa(42)")
	fmt.Println("  i, _ := strconv.Atoi(\"42\")")
	fmt.Println("  f, _ := strconv.ParseFloat(\"3.14\", 64)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Int to string: %s\n", s)
	fmt.Printf("  String to int: %d\n", i)
	fmt.Printf("  Parse float: %.2f\n", f)
	fmt.Println()
}

func main() {
	printTheory()
	exampleBytesBuffer()
	exampleStrconv()
}
