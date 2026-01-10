package main

import (
	_ "embed"
	"fmt"
)

//go:embed README.md
var readmeContent string

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Go Embed                                  ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📦 GO:EMBED")
	fmt.Println("   Embed files at compile time")
	fmt.Println("   //go:embed file.txt")
	fmt.Println("   var content string")
	fmt.Println()
}

func exampleEmbed() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Embed File")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  //go:embed README.md")
	fmt.Println("  var readmeContent string")
	fmt.Println()
	fmt.Println("Output:")
	if len(readmeContent) > 0 {
		fmt.Printf("  Embedded file length: %d bytes\n", len(readmeContent))
		fmt.Println("  (File content embedded at compile time)")
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleEmbed()
}
