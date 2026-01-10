package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: CGO                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔗 CGO")
	fmt.Println("   C interoperability")
	fmt.Println("   import \"C\"")
	fmt.Println()
}

func main() {
	printTheory()
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: CGO")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Code:")
	fmt.Println("  /*")
	fmt.Println("  #include <stdio.h>")
	fmt.Println("  void hello() { printf(\"Hello from C\\n\"); }")
	fmt.Println("  */")
	fmt.Println("  import \"C\"")
	fmt.Println("  C.hello()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  (CGO requires C compiler)")
	fmt.Println("  (This example doesn't use CGO)")
	fmt.Println()
}
