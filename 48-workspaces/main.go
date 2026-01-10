package main

import "fmt"

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Workspaces                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📁 WORKSPACES")
	fmt.Println("   go work init")
	fmt.Println("   go work use ./module1")
	fmt.Println()
}

func main() {
	printTheory()
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Workspace Setup")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Code:")
	fmt.Println("  go work init")
	fmt.Println("  go work use ./module1")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  Workspace initialized")
	fmt.Println("  (Run commands in workspace directory)")
	fmt.Println()
}
