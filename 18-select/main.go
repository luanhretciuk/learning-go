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
	fmt.Println("   • Like a switch statement for channels")
	fmt.Println("   • Blocks until at least one case is ready")
	fmt.Println("   • If multiple cases ready, one is chosen randomly")
	fmt.Println("   • Only one case executes per select")
	fmt.Println()
	fmt.Println("📋 SELECT SYNTAX")
	fmt.Println("   select {")
	fmt.Println("   case msg := <-ch1:     // receive from ch1")
	fmt.Println("       // handle msg")
	fmt.Println("   case ch2 <- value:     // send to ch2")
	fmt.Println("       // send succeeded")
	fmt.Println("   case <-time.After(1s): // timeout")
	fmt.Println("       // handle timeout")
	fmt.Println("   default:                // non-blocking")
	fmt.Println("       // no case ready")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("⚡ DEFAULT CASE")
	fmt.Println("   Makes select non-blocking")
	fmt.Println("   • Executes immediately if no case is ready")
	fmt.Println("   • Useful for checking channel availability")
	fmt.Println("   • Use with caution in loops (can cause CPU spinning)")
	fmt.Println()
	fmt.Println("⏱️  TIMEOUT PATTERN")
	fmt.Println("   Use time.After for timeouts")
	fmt.Println("   • Returns channel that receives after duration")
	fmt.Println("   • Common pattern for operation timeouts")
	fmt.Println("   • Prefer context.WithTimeout for better control")
	fmt.Println()
	fmt.Println("🔄 SELECT IN LOOPS")
	fmt.Println("   Common pattern for workers and servers")
	fmt.Println("   • Loop with select handles multiple channels")
	fmt.Println("   • Use done channel to exit gracefully")
	fmt.Println("   • Each iteration handles one event")
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
	fmt.Println("  ch1 := make(chan string, 1)")
	fmt.Println("  ch2 := make(chan string, 1)")
	fmt.Println("  ch1 <- \"message from ch1\"")
	fmt.Println("  select {")
	fmt.Println("  case msg1 := <-ch1: ...")
	fmt.Println("  case msg2 := <-ch2: ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • ch1 has a value, so first case is ready")
	fmt.Println("  • ch2 is empty, so second case blocks")
	fmt.Println("  • Select chooses the ready case (ch1)")
	fmt.Println("  • Only one case executes")
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

func exampleSelectMultipleReady() {
	fmt.Println("Example: Multiple Cases Ready")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "message 1"
	ch2 <- "message 2"

	fmt.Println("Code:")
	fmt.Println("  ch1 <- \"message 1\"")
	fmt.Println("  ch2 <- \"message 2\"")
	fmt.Println("  select { ... }  // both cases ready!")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Both channels have values")
	fmt.Println("  • Both cases are ready")
	fmt.Println("  • Go chooses ONE randomly (fairness)")
	fmt.Println("  • Order is non-deterministic")
	fmt.Println()

	fmt.Println("Output (may vary):")
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
	fmt.Println("  ch := make(chan string)  // empty, unbuffered")
	fmt.Println("  select {")
	fmt.Println("  case msg := <-ch: ...")
	fmt.Println("  default: ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Channel is empty, so receive would block")
	fmt.Println("  • default case executes immediately")
	fmt.Println("  • Select doesn't block")
	fmt.Println("  • Useful for checking availability")
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
	fmt.Println("  go func() { time.Sleep(200ms); ch <- \"message\" }()")
	fmt.Println("  select {")
	fmt.Println("  case msg := <-ch: ...")
	fmt.Println("  case <-time.After(100ms): ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Goroutine sends after 200ms")
	fmt.Println("  • Timeout fires after 100ms")
	fmt.Println("  • Timeout case executes first")
	fmt.Println("  • Operation times out before completion")
	fmt.Println()

	fmt.Println("Output:")
	select {
	case msg := <-ch:
		fmt.Printf("  Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  Timeout!")
	}
	time.Sleep(150 * time.Millisecond) // Let goroutine finish
	fmt.Println()
}

func exampleSelectInLoop() {
	fmt.Println("Example: Select in Loop (Worker Pattern)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	done := make(chan bool)

	// Send some messages
	go func() {
		ch1 <- "message 1"
		time.Sleep(50 * time.Millisecond)
		ch2 <- "message 2"
		time.Sleep(50 * time.Millisecond)
		ch1 <- "message 3"
		time.Sleep(50 * time.Millisecond)
		done <- true
	}()

	fmt.Println("Code:")
	fmt.Println("  for {")
	fmt.Println("      select {")
	fmt.Println("      case msg := <-ch1: ...")
	fmt.Println("      case msg := <-ch2: ...")
	fmt.Println("      case <-done: return")
	fmt.Println("      }")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Loop processes messages from multiple channels")
	fmt.Println("  • Each iteration handles one event")
	fmt.Println("  • done channel allows graceful exit")
	fmt.Println("  • Common pattern for workers and servers")
	fmt.Println()

	fmt.Println("Output:")
	count := 0
	for {
		select {
		case msg := <-ch1:
			fmt.Printf("  [ch1] %s\n", msg)
			count++
		case msg := <-ch2:
			fmt.Printf("  [ch2] %s\n", msg)
			count++
		case <-done:
			fmt.Println("  Done!")
			return
		}
		if count >= 3 {
			<-done // Consume done signal
			break
		}
	}
	fmt.Println()
}

func exampleSelectWithSend() {
	fmt.Println("Example: Select with Send Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	fmt.Println("Code:")
	fmt.Println("  select {")
	fmt.Println("  case ch1 <- 1: ...")
	fmt.Println("  case ch2 <- 2: ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Can select on send operations too")
	fmt.Println("  • Case executes when send succeeds")
	fmt.Println("  • Useful for non-blocking sends")
	fmt.Println()

	fmt.Println("Output:")
	select {
	case ch1 <- 1:
		fmt.Println("  Sent 1 to ch1")
	case ch2 <- 2:
		fmt.Println("  Sent 2 to ch2")
	}

	// Receive to show it worked
	fmt.Printf("  ch1 has: %d\n", <-ch1)
	fmt.Println()
}

func exampleSelectTimeoutSuccess() {
	fmt.Println("Example: Timeout Pattern - Success Case")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- "success!"
	}()

	fmt.Println("Code:")
	fmt.Println("  select {")
	fmt.Println("  case msg := <-ch: ...")
	fmt.Println("  case <-time.After(100ms): ...")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Goroutine sends after 50ms")
	fmt.Println("  • Timeout is 100ms")
	fmt.Println("  • Message arrives before timeout")
	fmt.Println("  • Success case executes")
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
	exampleSelectMultipleReady()
	exampleDefaultCase()
	exampleTimeout()
	exampleSelectInLoop()
	exampleSelectWithSend()
	exampleSelectTimeoutSuccess()
}
