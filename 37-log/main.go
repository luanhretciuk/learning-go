package main

import (
	"fmt"
	"log"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Log                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📝 LOGGING")
	fmt.Println("   log.Print(\"message\")")
	fmt.Println("   log.Println(\"message\")")
	fmt.Println("   log.Printf(\"format %s\", \"value\")")
	fmt.Println()
}

func exampleLog() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Logging")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  log.Print(\"Info message\")")
	fmt.Println("  log.Printf(\"Value: %d\", 42)")
	fmt.Println()
	fmt.Println("Output:")
	log.Print("Info message")
	log.Printf("Value: %d", 42)
	fmt.Println()
}

func main() {
	printTheory()
	exampleLog()
}
