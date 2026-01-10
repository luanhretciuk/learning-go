package main

import (
	"fmt"
	"sync"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Race Detector                            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔍 RACE DETECTOR")
	fmt.Println("   go run -race program.go")
	fmt.Println("   go test -race")
	fmt.Println()
	fmt.Println("⚠️  DATA RACES")
	fmt.Println("   Concurrent access without synchronization")
	fmt.Println()
}

func exampleRaceFixed() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Race Condition (Fixed with Mutex)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	
	wg.Add(2)
	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}()
	
	wg.Wait()
	
	fmt.Println("Code:")
	fmt.Println("  var mu sync.Mutex")
	fmt.Println("  mu.Lock()")
	fmt.Println("  counter++")
	fmt.Println("  mu.Unlock()")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Counter: %d (race-free)\n", counter)
	fmt.Println()
}

func main() {
	printTheory()
	exampleRaceFixed()
}
