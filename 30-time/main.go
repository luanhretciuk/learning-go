package main

import (
	"fmt"
	"time"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Time Package                              ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("⏰ TIME")
	fmt.Println("   Represents a point in time")
	fmt.Println("   now := time.Now()")
	fmt.Println()
	fmt.Println("⏱️  DURATION")
	fmt.Println("   Time intervals")
	fmt.Println("   d := 5 * time.Second")
	fmt.Println()
	fmt.Println("📅 FORMATTING")
	fmt.Println("   t.Format(\"2006-01-02 15:04:05\")")
	fmt.Println()
}

func exampleTime() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Time Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	now := time.Now()
	
	fmt.Println("Code:")
	fmt.Println("  now := time.Now()")
	fmt.Println("  now.Format(\"2006-01-02 15:04:05\")")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Now: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Println()
}

func exampleDuration() {
	fmt.Println("Example: Duration")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	d1 := 5 * time.Second
	d2 := 10 * time.Minute
	
	fmt.Println("Code:")
	fmt.Println("  d1 := 5 * time.Second")
	fmt.Println("  d2 := 10 * time.Minute")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Duration 1: %v\n", d1)
	fmt.Printf("  Duration 2: %v\n", d2)
	fmt.Println()
}

func main() {
	printTheory()
	exampleTime()
	exampleDuration()
}
