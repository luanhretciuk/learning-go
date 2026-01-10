package main

import "fmt"

type Status int

const (
	Pending Status = iota
	Running
	Completed
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Go Generate                               ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔧 GO:GENERATE")
	fmt.Println("   //go:generate stringer -type=Status")
	fmt.Println("   Run: go generate")
	fmt.Println()
}

func (s Status) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Running:
		return "Running"
	case Completed:
		return "Completed"
	default:
		return "Unknown"
	}
}

func exampleGenerate() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Go Generate")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  //go:generate stringer -type=Status")
	fmt.Println("  type Status int")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Status: %s\n", Running)
	fmt.Println()
}

func main() {
	printTheory()
	exampleGenerate()
}
