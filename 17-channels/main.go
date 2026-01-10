package main

import (
	"fmt"
	"time"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Channels                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📡 CHANNELS")
	fmt.Println("   Go's way to communicate between goroutines")
	fmt.Println("   ch := make(chan int)")
	fmt.Println()
	fmt.Println("📤 SENDING AND RECEIVING")
	fmt.Println("   ch <- 42        // send")
	fmt.Println("   value := <-ch   // receive")
	fmt.Println()
	fmt.Println("🔄 UNBUFFERED CHANNELS")
	fmt.Println("   • Synchronous communication")
	fmt.Println("   • Sender blocks until receiver ready")
	fmt.Println()
	fmt.Println("📦 BUFFERED CHANNELS")
	fmt.Println("   • Asynchronous up to capacity")
	fmt.Println("   • ch := make(chan int, 10)")
	fmt.Println()
	fmt.Println("🚪 CLOSING CHANNELS")
	fmt.Println("   close(ch)")
	fmt.Println("   value, ok := <-ch  // ok is false when closed")
	fmt.Println()
}

func exampleUnbufferedChannel() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Unbuffered Channel")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ch := make(chan int)
	
	go func() {
		ch <- 42
		fmt.Println("  Sent: 42")
	}()
	
	fmt.Println("Code:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  go func() { ch <- 42 }()")
	fmt.Println("  value := <-ch")
	fmt.Println()
	
	value := <-ch
	fmt.Println("Output:")
	fmt.Printf("  Received: %d\n", value)
	fmt.Println()
}

func exampleBufferedChannel() {
	fmt.Println("Example: Buffered Channel")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ch := make(chan int, 3)
	
	ch <- 1
	ch <- 2
	ch <- 3
	
	fmt.Println("Code:")
	fmt.Println("  ch := make(chan int, 3)")
	fmt.Println("  ch <- 1")
	fmt.Println("  ch <- 2")
	fmt.Println("  ch <- 3")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Received: %d\n", <-ch)
	fmt.Printf("  Received: %d\n", <-ch)
	fmt.Printf("  Received: %d\n", <-ch)
	fmt.Println()
}

func exampleChannelClose() {
	fmt.Println("Example: Closing Channels")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	
	fmt.Println("Code:")
	fmt.Println("  ch := make(chan int, 2)")
	fmt.Println("  ch <- 1")
	fmt.Println("  ch <- 2")
	fmt.Println("  close(ch)")
	fmt.Println()
	fmt.Println("Output:")
	for i := 0; i < 3; i++ {
		value, ok := <-ch
		if ok {
			fmt.Printf("  Received: %d (ok: %t)\n", value, ok)
		} else {
			fmt.Printf("  Channel closed (ok: %t)\n", ok)
		}
	}
	fmt.Println()
}

func exampleChannelRange() {
	fmt.Println("Example: Range over Channel")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)
	
	fmt.Println("Code:")
	fmt.Println("  for value := range ch {")
	fmt.Println("      fmt.Println(value)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Output:")
	for value := range ch {
		fmt.Printf("  Received: %d\n", value)
	}
	fmt.Println()
}

func exampleChannelDirection() {
	fmt.Println("Example: Channel Directions")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	sendOnly := func(ch chan<- int) {
		ch <- 42
	}
	
	receiveOnly := func(ch <-chan int) int {
		return <-ch
	}
	
	ch := make(chan int, 1)
	sendOnly(ch)
	value := receiveOnly(ch)
	
	fmt.Println("Code:")
	fmt.Println("  func sendOnly(ch chan<- int) { ch <- 42 }")
	fmt.Println("  func receiveOnly(ch <-chan int) int { return <-ch }")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Received: %d\n", value)
	fmt.Println()
}

func main() {
	printTheory()
	exampleUnbufferedChannel()
	exampleBufferedChannel()
	exampleChannelClose()
	exampleChannelRange()
	exampleChannelDirection()
}
