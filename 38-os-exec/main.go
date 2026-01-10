package main

import (
	"fmt"
	"os/exec"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: OS Exec                                   ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("⚙️  EXECUTING COMMANDS")
	fmt.Println("   cmd := exec.Command(\"ls\", \"-l\")")
	fmt.Println("   output, err := cmd.Output()")
	fmt.Println()
}

func exampleExec() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Execute Command")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	fmt.Println("Code:")
	fmt.Println("  cmd := exec.Command(\"echo\", \"Hello\")")
	fmt.Println("  output, err := cmd.Output()")
	fmt.Println()
	
	cmd := exec.Command("echo", "Hello")
	output, err := cmd.Output()
	
	fmt.Println("Output:")
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  %s", string(output))
	}
	fmt.Println()
}

func main() {
	printTheory()
	exampleExec()
}
