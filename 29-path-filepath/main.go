package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Path & Filepath                            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("📁 PATH PACKAGE")
	fmt.Println("   URL/path manipulation")
	fmt.Println()
	fmt.Println("📂 FILEPATH PACKAGE")
	fmt.Println("   OS-specific path handling")
	fmt.Println()
}

func examplePath() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Path Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	p := "/usr/bin/go"
	
	fmt.Println("Code:")
	fmt.Println("  path.Base(\"/usr/bin/go\")")
	fmt.Println("  path.Dir(\"/usr/bin/go\")")
	fmt.Println("  path.Ext(\"file.txt\")")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Base: %s\n", path.Base(p))
	fmt.Printf("  Dir: %s\n", path.Dir(p))
	fmt.Printf("  Ext: %s\n", path.Ext("file.txt"))
	fmt.Println()
}

func exampleFilepath() {
	fmt.Println("Example: Filepath Operations")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	joined := filepath.Join("dir", "subdir", "file.txt")
	
	fmt.Println("Code:")
	fmt.Println("  filepath.Join(\"dir\", \"subdir\", \"file.txt\")")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Joined: %s\n", joined)
	fmt.Println()
}

func main() {
	printTheory()
	examplePath()
	exampleFilepath()
}
