package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func printTheory() {
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║          TOPIC: Encoding                                 ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("📚 THEORY")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("🔢 BASE64")
	fmt.Println("   Base64 encoding/decoding")
	fmt.Println("   base64.StdEncoding.EncodeToString(data)")
	fmt.Println()
	fmt.Println("🔢 HEX")
	fmt.Println("   Hexadecimal encoding")
	fmt.Println("   hex.EncodeToString(data)")
	fmt.Println()
}

func exampleBase64() {
	fmt.Println("🔧 PRACTICAL EXAMPLES")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	fmt.Println("Example: Base64 Encoding")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	data := []byte("Hello, World!")
	encoded := base64.StdEncoding.EncodeToString(data)
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	
	fmt.Println("Code:")
	fmt.Println("  encoded := base64.StdEncoding.EncodeToString(data)")
	fmt.Println("  decoded, _ := base64.StdEncoding.DecodeString(encoded)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Original: %s\n", string(data))
	fmt.Printf("  Encoded: %s\n", encoded)
	fmt.Printf("  Decoded: %s\n", string(decoded))
	fmt.Println()
}

func exampleHex() {
	fmt.Println("Example: Hex Encoding")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
	
	data := []byte("Hello")
	encoded := hex.EncodeToString(data)
	decoded, _ := hex.DecodeString(encoded)
	
	fmt.Println("Code:")
	fmt.Println("  encoded := hex.EncodeToString(data)")
	fmt.Println("  decoded, _ := hex.DecodeString(encoded)")
	fmt.Println()
	fmt.Println("Output:")
	fmt.Printf("  Original: %s\n", string(data))
	fmt.Printf("  Encoded: %s\n", encoded)
	fmt.Printf("  Decoded: %s\n", string(decoded))
	fmt.Println()
}

func main() {
	printTheory()
	exampleBase64()
	exampleHex()
}
