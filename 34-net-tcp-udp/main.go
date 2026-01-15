package main

import (
	"fmt"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Net TCP/UDP                               ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🌐 TCP/UDP")
	fmt.Println("   Low-level networking")
	fmt.Println("   ln, _ := net.Listen(\"tcp\", \":8080\")")
	fmt.Println("   conn, _ := net.Dial(\"tcp\", \"localhost:8080\")")
	fmt.Println()
}

func exampleTCP() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: TCP Connection")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Code:")
	fmt.Println("  ln, _ := net.Listen(\"tcp\", \":8080\")")
	fmt.Println("  conn, _ := ln.Accept()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  TCP server would listen on :8080")
	fmt.Println("  (Not started in this example)")
	fmt.Println()
}

func main() {
	printTheory()
	exampleTCP()
}
