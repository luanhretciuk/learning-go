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
	fmt.Println("   Provides cancellation, timeouts, and request-scoped values")
	fmt.Println("   • Standard way to handle cancellation in Go")
	fmt.Println("   • Propagates through call chains")
	fmt.Println("   • Immutable (each operation returns new context)")
	fmt.Println("   • First parameter in functions (ctx context.Context)")
	fmt.Println()
	fmt.Println("🚫 CANCELLATION")
	fmt.Println("   ctx, cancel := context.WithCancel(parent)")
	fmt.Println("   cancel()  // signal cancellation")
	fmt.Println("   • Cancels context and all children")
	fmt.Println("   • Check with ctx.Done() channel")
	fmt.Println("   • Always call cancel() to free resources")
	fmt.Println()
	fmt.Println("⏱️  TIMEOUT")
	fmt.Println("   ctx, cancel := context.WithTimeout(parent, duration)")
	fmt.Println("   • Automatically cancels after duration")
	fmt.Println("   • Useful for operation timeouts")
	fmt.Println("   • Still need to call cancel() for cleanup")
	fmt.Println()
	fmt.Println("⏰ DEADLINE")
	fmt.Println("   ctx, cancel := context.WithDeadline(parent, time)")
	fmt.Println("   • Cancels at specific time")
	fmt.Println("   • Similar to timeout but absolute time")
	fmt.Println()
	fmt.Println("📝 VALUES")
	fmt.Println("   ctx := context.WithValue(parent, key, value)")
	fmt.Println("   • Request-scoped values")
	fmt.Println("   • Use sparingly (not for optional parameters)")
	fmt.Println("   • Prefer explicit parameters")
	fmt.Println()
	fmt.Println("🔍 CONTEXT METHODS")
	fmt.Println("   ctx.Done()        // channel that closes on cancellation")
	fmt.Println("   ctx.Err()         // error (context.Canceled or context.DeadlineExceeded)")
	fmt.Println("   ctx.Value(key)    // get value from context")
	fmt.Println("   ctx.Deadline()    // deadline time (if set)")
	fmt.Println()
}

func exampleWithCancel() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: WithCancel - Manual Cancellation")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  [Goroutine] Cancelling context...")
		cancel()
	}()

	fmt.Println("Code:")
	fmt.Println("  ctx, cancel := context.WithCancel(context.Background())")
	fmt.Println("  go func() { cancel() }()")
	fmt.Println("  select {")
	fmt.Println("  case <-ctx.Done(): ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • WithCancel returns context and cancel function")
	fmt.Println("  • Calling cancel() signals cancellation")
	fmt.Println("  • ctx.Done() channel closes when cancelled")
	fmt.Println("  • All child contexts are also cancelled")
	fmt.Println()

	fmt.Println("Output:")
	select {
	case <-ctx.Done():
		fmt.Printf("  Context cancelled: %v\n", ctx.Err())
	case <-time.After(200 * time.Millisecond):
		fmt.Println("  Timeout (shouldn't happen)")
	}
	fmt.Println()
}

func exampleWithTimeout() {
	fmt.Println("Example: WithTimeout - Automatic Timeout")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel() // Always call cancel to free resources

	fmt.Println("Code:")
	fmt.Println("  ctx, cancel := context.WithTimeout(parent, 50*time.Millisecond)")
	fmt.Println("  defer cancel()  // cleanup")
	fmt.Println("  select {")
	fmt.Println("  case <-ctx.Done(): ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Context automatically cancels after duration")
	fmt.Println("  • Still need to call cancel() for cleanup")
	fmt.Println("  • Useful for operation timeouts")
	fmt.Println("  • Better than time.After() in select")
	fmt.Println()

	fmt.Println("Output:")
	select {
	case <-ctx.Done():
		fmt.Printf("  Context done: %v\n", ctx.Err())
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  Work completed (shouldn't happen)")
	}
	fmt.Println()
}

func exampleWithDeadline() {
	fmt.Println("Example: WithDeadline - Absolute Time")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	deadline := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Println("Code:")
	fmt.Println("  deadline := time.Now().Add(100 * time.Millisecond)")
	fmt.Println("  ctx, cancel := context.WithDeadline(parent, deadline)")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Similar to WithTimeout but uses absolute time")
	fmt.Println("  • Useful when you have a specific deadline")
	fmt.Println("  • Cancels at the specified time")
	fmt.Println()

	fmt.Println("Output:")
	select {
	case <-ctx.Done():
		fmt.Printf("  Context done: %v\n", ctx.Err())
		if deadline, ok := ctx.Deadline(); ok {
			fmt.Printf("  Deadline was: %v\n", deadline)
		}
	case <-time.After(150 * time.Millisecond):
		fmt.Println("  Timeout (shouldn't happen)")
	}
	fmt.Println()
}

func exampleWithValue() {
	fmt.Println("Example: WithValue - Request-Scoped Values")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	type key string
	userIDKey := key("userID")
	requestIDKey := key("requestID")

	ctx := context.WithValue(context.Background(), userIDKey, 123)
	ctx = context.WithValue(ctx, requestIDKey, "req-456")

	fmt.Println("Code:")
	fmt.Println("  type key string")
	fmt.Println("  ctx := context.WithValue(parent, \"userID\", 123)")
	fmt.Println("  value := ctx.Value(\"userID\")")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Store request-scoped values in context")
	fmt.Println("  • Use custom types for keys (avoid collisions)")
	fmt.Println("  • Use sparingly - prefer explicit parameters")
	fmt.Println("  • Common for: user ID, request ID, trace ID")
	fmt.Println()

	fmt.Println("Output:")
	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)
	fmt.Printf("  userID: %v\n", userID)
	fmt.Printf("  requestID: %v\n", requestID)
	fmt.Println()
}

func exampleContextPropagation() {
	fmt.Println("Example: Context Propagation")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	work := func(ctx context.Context, name string) {
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  %s: cancelled\n", name)
				return
			default:
				fmt.Printf("  %s: working %d\n", name, i)
				time.Sleep(50 * time.Millisecond)
			}
		}
		fmt.Printf("  %s: completed\n", name)
	}

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Code:")
	fmt.Println("  ctx, cancel := context.WithCancel(parent)")
	fmt.Println("  go work(ctx, \"worker1\")")
	fmt.Println("  go work(ctx, \"worker2\")")
	fmt.Println("  cancel()  // cancels all workers")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Pass context to all goroutines")
	fmt.Println("  • Cancelling parent cancels all children")
	fmt.Println("  • Workers check ctx.Done() to stop gracefully")
	fmt.Println()

	fmt.Println("Output:")
	go work(ctx, "worker1")
	go work(ctx, "worker2")

	time.Sleep(150 * time.Millisecond)
	fmt.Println("  [Main] Cancelling...")
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func exampleHTTPTimeout() {
	fmt.Println("Example: HTTP Request with Timeout")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Code (typical HTTP pattern):")
	fmt.Println("  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)")
	fmt.Println("  defer cancel()")
	fmt.Println("  req, _ := http.NewRequestWithContext(ctx, \"GET\", url, nil)")
	fmt.Println("  resp, err := http.DefaultClient.Do(req)")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Set timeout for HTTP requests")
	fmt.Println("  • Request automatically cancelled on timeout")
	fmt.Println("  • Prevents hanging requests")
	fmt.Println("  • Standard pattern in Go HTTP clients")
	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	fmt.Println("Output (simulated):")
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("  Request would succeed")
	case <-ctx.Done():
		fmt.Printf("  Request timeout: %v\n", ctx.Err())
	}
	fmt.Println()
}

func exampleContextInFunction() {
	fmt.Println("Example: Context as First Parameter")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	processData := func(ctx context.Context, data string) error {
		// Check if cancelled before starting
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// Simulate work with cancellation check
		for i := 0; i < 3; i++ {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				fmt.Printf("  Processing %s: step %d\n", data, i)
				time.Sleep(50 * time.Millisecond)
			}
		}
		return nil
	}

	fmt.Println("Code:")
	fmt.Println("  func processData(ctx context.Context, data string) error {")
	fmt.Println("      select {")
	fmt.Println("      case <-ctx.Done():")
	fmt.Println("          return ctx.Err()")
	fmt.Println("      default:")
	fmt.Println("          // do work")
	fmt.Println("      }")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Context is always first parameter")
	fmt.Println("  • Check ctx.Done() in loops")
	fmt.Println("  • Return ctx.Err() when cancelled")
	fmt.Println("  • Allows caller to cancel operation")
	fmt.Println()

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Output:")
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	err := processData(ctx, "data1")
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	}
	fmt.Println()
}

func exampleBackgroundContext() {
	fmt.Println("Example: Background and TODO Contexts")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Code:")
	fmt.Println("  ctx1 := context.Background()  // root context, never cancelled")
	fmt.Println("  ctx2 := context.TODO()       // placeholder, should be replaced")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • context.Background(): root context, never cancelled")
	fmt.Println("  • Use as parent for top-level contexts")
	fmt.Println("  • context.TODO(): placeholder when context is unclear")
	fmt.Println("  • Both return empty contexts")
	fmt.Println()

	bgCtx := context.Background()
	todoCtx := context.TODO()

	fmt.Println("Output:")
	fmt.Printf("  Background context: %v\n", bgCtx)
	fmt.Printf("  TODO context: %v\n", todoCtx)
	fmt.Printf("  Background == TODO: %t\n", bgCtx == todoCtx)
	fmt.Println()
}

func main() {
	printTheory()
	exampleWithCancel()
	exampleWithTimeout()
	exampleWithDeadline()
	exampleWithValue()
	exampleContextPropagation()
	exampleHTTPTimeout()
	exampleContextInFunction()
	exampleBackgroundContext()
}
