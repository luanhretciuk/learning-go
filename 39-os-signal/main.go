package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: OS Signal                                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📡 SIGNAL HANDLING")
	fmt.Println("   c := make(chan os.Signal, 1)")
	fmt.Println("   signal.Notify(c, os.Interrupt)")
	fmt.Println()
}

func exampleSignal() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Signal Handling")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	fmt.Println("Code:")
	fmt.Println("  c := make(chan os.Signal, 1)")
	fmt.Println("  signal.Notify(c, os.Interrupt)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Println("  Signal handler registered")
	fmt.Println("  (Would catch SIGINT/SIGTERM in real usage)")
	fmt.Println()
	
	// Don't block - just show the pattern
	select {
	case sig := <-c:
		fmt.Printf("  Received signal: %v\n", sig)
	default:
		// No signal received
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleSignal()
}
