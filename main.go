package main

import (
	"fmt"
	"os"

	"github.com/storozhenko98/Q/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: your-cli-tool <command> [name]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "hello":
		name := "World"
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		fmt.Println(utils.Greet(name))
	case "version":
		fmt.Println("v0.1.0")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
