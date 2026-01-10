package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Build Tags                                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🏷️  BUILD TAGS")
	fmt.Println("   //go:build linux")
	fmt.Println("   // +build linux")
	fmt.Println()
}

func main() {
	printTheory()
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Build Tags")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Code:")
	fmt.Println("  //go:build linux")
	fmt.Println("  package main")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  This file would only compile on Linux")
	fmt.Println()
}
