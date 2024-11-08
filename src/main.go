package main

import (
	"cli-file-tool/src/internal/fileops"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-file-tool <command> [args...]")
		return
	}
	command := os.Args[1]

	switch command {
	case "list":
		//initialize flags
		var dirOnly, fileOnly bool

		//check flags
		for _, arg := range os.Args[2:] {
			switch arg {
			case "--dir-only":
				dirOnly = true
			case "--file-only":
				fileOnly = true
			}

		}
		fileops.ListFiles(dirOnly, fileOnly)
	case "help":
		showHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}

}

func showHelp() {
	fmt.Println("Usage: cli-file-tool <command> [args...]")
	fmt.Println("Available commands:")
	fmt.Println("  list    List files in the current directory")
	fmt.Println("          Flags:")
	fmt.Println("            --verbose   Show detailed output")
	fmt.Println("  help    Show this help message")
}
