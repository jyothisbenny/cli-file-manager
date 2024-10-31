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

		listFiles(dirOnly, fileOnly)
	case "help":
		showHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}

}

func listFiles(dirOnly, fileOnly bool) {
	files, err := fileops.ListFile(dirOnly, fileOnly)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {

		info, err := os.Stat(file)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if info.IsDir() {
			fmt.Printf("\033[34m%s\033[0m\n", file)
		} else {
			fmt.Printf("\033[32m%s\033[0m\n", file)
		}
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
