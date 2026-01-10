package main

import (
	"fmt"
	"learning-go/23-packages/utils" // This would be a local package in real usage
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Packages                                  ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📦 PACKAGES")
	fmt.Println("   Organize Go code into reusable units")
	fmt.Println("   package main  // executable")
	fmt.Println("   package utils // library")
	fmt.Println()
	fmt.Println("📤 EXPORTS")
	fmt.Println("   • Capitalized = exported (visible outside)")
	fmt.Println("   • Lowercase = unexported (package-private)")
	fmt.Println()
	fmt.Println("📥 IMPORTS")
	fmt.Println("   import \"fmt\"")
	fmt.Println("   import (")
	fmt.Println("       \"fmt\"")
	fmt.Println("       \"os\"")
	fmt.Println("   )")
	fmt.Println()
	fmt.Println("🚀 INIT FUNCTIONS")
	fmt.Println("   func init() { ... }")
	fmt.Println("   • Run automatically on package import")
	fmt.Println("   • Multiple init functions allowed")
	fmt.Println()
}

func examplePackageUsage() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Using Packages")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  import \"learning-go/23-packages/utils\"")
	fmt.Println("  result := utils.Add(5, 3)")
	fmt.Println()
	
	result := utils.Add(5, 3)
	fmt.Println("Output:")
	fmt.Printf("  Result: %d\n", result)
	fmt.Println()
}

func main() {
	printTheory()
	examplePackageUsage()
}
