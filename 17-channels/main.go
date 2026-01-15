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
	fmt.Println("   • Type-safe communication")
	fmt.Println("   • Synchronization built-in")
	fmt.Println("   • FIFO queue behavior")
	fmt.Println("   • Syntax: ch := make(chan Type)")
	fmt.Println()
	fmt.Println("📤 SENDING AND RECEIVING")
	fmt.Println("   ch <- 42        // send value to channel")
	fmt.Println("   value := <-ch   // receive value from channel")
	fmt.Println("   <-ch            // receive and discard")
	fmt.Println()
	fmt.Println("🔄 UNBUFFERED CHANNELS")
	fmt.Println("   ch := make(chan int)  // no capacity specified")
	fmt.Println("   • Synchronous communication")
	fmt.Println("   • Sender blocks until receiver is ready")
	fmt.Println("   • Receiver blocks until sender is ready")
	fmt.Println("   • Direct handoff - no storage")
	fmt.Println("   • Guarantees receiver gets value")
	fmt.Println()
	fmt.Println("📦 BUFFERED CHANNELS")
	fmt.Println("   ch := make(chan int, 10)  // capacity 10")
	fmt.Println("   • Asynchronous up to capacity")
	fmt.Println("   • Stores values in FIFO queue")
	fmt.Println("   • Sender blocks only when buffer is full")
	fmt.Println("   • Receiver blocks only when buffer is empty")
	fmt.Println("   • Faster for bursty workloads")
	fmt.Println()
	fmt.Println("🚪 CLOSING CHANNELS")
	fmt.Println("   close(ch)  // only sender should close")
	fmt.Println("   • Sending to closed channel = panic")
	fmt.Println("   • Receiving from closed channel = zero value + ok=false")
	fmt.Println("   • Closing already closed channel = panic")
	fmt.Println("   • Use to signal 'no more values'")
	fmt.Println()
	fmt.Println("🔄 CHANNEL DIRECTIONS")
	fmt.Println("   chan<- T    // send-only channel")
	fmt.Println("   <-chan T    // receive-only channel")
	fmt.Println("   chan T      // bidirectional (default)")
	fmt.Println("   • Enforces communication patterns")
	fmt.Println("   • Prevents accidental misuse")
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
		fmt.Println("  [Goroutine] About to send 42...")
		ch <- 42  // Blocks here until receiver is ready
		fmt.Println("  [Goroutine] Sent 42!")
	}()

	time.Sleep(50 * time.Millisecond) // Give goroutine time to start

	fmt.Println("Code:")
	fmt.Println("  ch := make(chan int)  // unbuffered")
	fmt.Println("  go func() { ch <- 42 }()")
	fmt.Println("  value := <-ch")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Sender blocks until receiver reads")
	fmt.Println("  • This is synchronous communication")
	fmt.Println("  • Guarantees the value is received")
	fmt.Println()

	fmt.Println("Output:")
	fmt.Println("  [Main] About to receive...")
	value := <-ch  // Unblocks sender
	fmt.Printf("  [Main] Received: %d\n", value)
	time.Sleep(50 * time.Millisecond) // Wait for goroutine to finish
	fmt.Println()
}

func exampleBufferedChannel() {
	fmt.Println("Example: Buffered Channel")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	ch := make(chan int, 3)

	fmt.Println("Code:")
	fmt.Println("  ch := make(chan int, 3)  // buffer capacity 3")
	fmt.Println("  ch <- 1  // doesn't block")
	fmt.Println("  ch <- 2  // doesn't block")
	fmt.Println("  ch <- 3  // doesn't block")
	fmt.Println("  ch <- 4  // would block (buffer full)")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Can send up to 3 values without blocking")
	fmt.Println("  • Values stored in FIFO queue")
	fmt.Println("  • Receiver gets values in order sent")
	fmt.Println()

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Output:")
	fmt.Println("  Sent 3 values to buffered channel")
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
	fmt.Println("  value, ok := <-ch  // ok=false when closed")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • close() signals 'no more values'")
	fmt.Println("  • Can still receive remaining values")
	fmt.Println("  • After all values received, ok=false")
	fmt.Println("  • Zero value returned when closed and empty")
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
	close(ch)  // Must close before ranging!

	fmt.Println("Code:")
	fmt.Println("  ch := make(chan int, 3)")
	fmt.Println("  ch <- 10; ch <- 20; ch <- 30")
	fmt.Println("  close(ch)  // IMPORTANT: close before range")
	fmt.Println("  for value := range ch {")
	fmt.Println("      fmt.Println(value)")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • range automatically receives until channel closes")
	fmt.Println("  • Loop exits when channel is closed and empty")
	fmt.Println("  • Must close channel or range will block forever")
	fmt.Println()

	fmt.Println("Output:")
	for value := range ch {
		fmt.Printf("  Received: %d\n", value)
	}
	fmt.Println("  Loop exited (channel closed)")
	fmt.Println()
}

func exampleChannelDirection() {
	fmt.Println("Example: Channel Directions")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	sendOnly := func(ch chan<- int) {
		ch <- 42
		// <-ch  // ERROR: cannot receive from send-only channel
	}

	receiveOnly := func(ch <-chan int) int {
		return <-ch
		// ch <- 42  // ERROR: cannot send to receive-only channel
	}

	ch := make(chan int, 1)
	sendOnly(ch)
	value := receiveOnly(ch)

	fmt.Println("Code:")
	fmt.Println("  func sendOnly(ch chan<- int) { ch <- 42 }")
	fmt.Println("  func receiveOnly(ch <-chan int) int { return <-ch }")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • chan<- T: can only send (send-only)")
	fmt.Println("  • <-chan T: can only receive (receive-only)")
	fmt.Println("  • Enforces communication patterns at compile time")
	fmt.Println("  • Prevents accidental misuse")
	fmt.Println()

	fmt.Println("Output:")
	fmt.Printf("  Received: %d\n", value)
	fmt.Println()
}

func exampleChannelFIFO() {
	fmt.Println("Example: Channel FIFO Behavior")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	ch := make(chan int, 5)

	fmt.Println("Code:")
	fmt.Println("  ch := make(chan int, 5)")
	fmt.Println("  ch <- 1; ch <- 2; ch <- 3; ch <- 4; ch <- 5")
	fmt.Println("  // Receive in order: 1, 2, 3, 4, 5")
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println("  • Channels are FIFO (First In, First Out)")
	fmt.Println("  • Values received in order they were sent")
	fmt.Println("  • Like a queue data structure")
	fmt.Println()

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	fmt.Println("Output:")
	fmt.Println("  Sent: 1, 2, 3, 4, 5")
	fmt.Println("  Receiving:")
	for i := 0; i < 5; i++ {
		fmt.Printf("    %d\n", <-ch)
	}
	fmt.Println()
}

func exampleChannelBlocking() {
	fmt.Println("Example: Channel Blocking Behavior")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	fmt.Println("Unbuffered channel blocking:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  ch <- 1  // BLOCKS until receiver ready")
	fmt.Println()

	fmt.Println("Buffered channel (not full):")
	fmt.Println("  ch := make(chan int, 2)")
	fmt.Println("  ch <- 1  // doesn't block (buffer has space)")
	fmt.Println("  ch <- 2  // doesn't block (buffer has space)")
	fmt.Println("  ch <- 3  // BLOCKS (buffer full)")
	fmt.Println()

	fmt.Println("Receiving from empty channel:")
	fmt.Println("  ch := make(chan int)")
	fmt.Println("  value := <-ch  // BLOCKS until sender ready")
	fmt.Println()

	fmt.Println("Receiving from buffered channel:")
	fmt.Println("  ch := make(chan int, 2)")
	fmt.Println("  ch <- 1; ch <- 2")
	fmt.Println("  value := <-ch  // doesn't block (buffer has values)")
	fmt.Println()
}

func main() {
	printTheory()
	exampleUnbufferedChannel()
	exampleBufferedChannel()
	exampleChannelClose()
	exampleChannelRange()
	exampleChannelDirection()
	exampleChannelFIFO()
	exampleChannelBlocking()
}
