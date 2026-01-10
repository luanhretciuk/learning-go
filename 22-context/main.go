package main

import (
	"context"
	"fmt"
	"time"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Context Package                           ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📦 CONTEXT")
	fmt.Println("   Provides cancellation, timeouts, and values")
	fmt.Println("   ctx := context.Background()")
	fmt.Println()
	fmt.Println("🚫 CANCELLATION")
	fmt.Println("   ctx, cancel := context.WithCancel(parent)")
	fmt.Println("   cancel()  // signal cancellation")
	fmt.Println()
	fmt.Println("⏱️  TIMEOUT")
	fmt.Println("   ctx, cancel := context.WithTimeout(parent, 5*time.Second)")
	fmt.Println()
	fmt.Println("📝 VALUES")
	fmt.Println("   ctx := context.WithValue(parent, \"key\", \"value\")")
	fmt.Println()
}

func exampleWithCancel() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: With Cancel")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	
	fmt.Println("Code:")
	fmt.Println("  ctx, cancel := context.WithCancel(context.Background())")
	fmt.Println("  cancel()  // signal cancellation")
	fmt.Println()
	
	select {
	case <-ctx.Done():
		fmt.Println("Output:")
		fmt.Printf("  Context cancelled: %v\n", ctx.Err())
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Output:")
		fmt.Println("  Timeout")
	}
	fmt.Println()
}

func exampleWithTimeout() {
	fmt.Println("Example: With Timeout")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	
	fmt.Println("Code:")
	fmt.Println("  ctx, cancel := context.WithTimeout(parent, 50*time.Millisecond)")
	fmt.Println()
	
	select {
	case <-ctx.Done():
		fmt.Println("Output:")
		fmt.Printf("  Context done: %v\n", ctx.Err())
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Output:")
		fmt.Println("  Work completed")
	}
	fmt.Println()
}

func exampleWithValue() {
	fmt.Println("Example: With Value")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ctx := context.WithValue(context.Background(), "userID", 123)
	
	fmt.Println("Code:")
	fmt.Println("  ctx := context.WithValue(parent, \"userID\", 123)")
	fmt.Println("  value := ctx.Value(\"userID\")")
	fmt.Println()
	
	value := ctx.Value("userID")
	fmt.Println("Output:")
	fmt.Printf("  userID: %v\n", value)
	fmt.Println()
}

func main() {
	printTheory()
	exampleWithCancel()
	exampleWithTimeout()
	exampleWithValue()
}
