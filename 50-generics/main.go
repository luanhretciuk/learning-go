package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Generics                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔢 GENERICS (Go 1.18+)")
	fmt.Println("   Type parameters")
	fmt.Println("   func Print[T any](v T) { ... }")
	fmt.Println()
}

func Print[T any](v T) {
	fmt.Println(v)
}

func exampleGenericFunction() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Generic Function")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  func Print[T any](v T) { fmt.Println(v) }")
	fmt.Println("  Print(42)")
	fmt.Println("  Print(\"hello\")")
	fmt.Println()
	fmt.Println("Output:")
	Print(42)
	Print("hello")
	fmt.Println()
}

func main() {
	printTheory()
	exampleGenericFunction()
}
