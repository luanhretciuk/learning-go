package main

import (
	"fmt"
	"time"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Select Statement                         ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔀 SELECT STATEMENT")
	fmt.Println("   Lets a goroutine wait on multiple channel operations")
	fmt.Println("   select {")
	fmt.Println("   case msg := <-ch1: ...")
	fmt.Println("   case msg := <-ch2: ...")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("⚡ DEFAULT CASE")
	fmt.Println("   Non-blocking select:")
	fmt.Println("   select {")
	fmt.Println("   case msg := <-ch: ...")
	fmt.Println("   default: ...")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("⏱️  TIMEOUT PATTERN")
	fmt.Println("   Use time.After for timeouts")
	fmt.Println()
}

func exampleBasicSelect() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Basic Select")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	
	ch1 <- "message from ch1"
	
	fmt.Println("Code:")
	fmt.Println("  select {")
	fmt.Println("  case msg1 := <-ch1: ...")
	fmt.Println("  case msg2 := <-ch2: ...")
	fmt.Println("  }")
	fmt.Println()
	
	fmt.Println("Output:")
	select {
	case msg1 := <-ch1:
		fmt.Printf("  Received from ch1: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("  Received from ch2: %s\n", msg2)
	}
	fmt.Println()
}

func exampleDefaultCase() {
	fmt.Println("Example: Default Case (Non-blocking)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ch := make(chan string)
	
	fmt.Println("Code:")
	fmt.Println("  select {")
	fmt.Println("  case msg := <-ch: ...")
	fmt.Println("  default: ...")
	fmt.Println("  }")
	fmt.Println()
	
	fmt.Println("Output:")
	select {
	case msg := <-ch:
		fmt.Printf("  Received: %s\n", msg)
	default:
		fmt.Println("  No message available (non-blocking)")
	}
	fmt.Println()
}

func exampleTimeout() {
	fmt.Println("Example: Timeout Pattern")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ch := make(chan string)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "message"
	}()
	
	fmt.Println("Code:")
	fmt.Println("  select {")
	fmt.Println("  case msg := <-ch: ...")
	fmt.Println("  case <-time.After(100 * time.Millisecond): ...")
	fmt.Println("  }")
	fmt.Println()
	
	fmt.Println("Output:")
	select {
	case msg := <-ch:
		fmt.Printf("  Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  Timeout!")
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleBasicSelect()
	exampleDefaultCase()
	exampleTimeout()
}
