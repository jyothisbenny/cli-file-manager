package main

import (
	"cli-file-tool/src/internal/fileops" // import fileops with full path
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
		files, err := fileops.ListFile()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, file := range files {
			fmt.Println(file)
		}
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
