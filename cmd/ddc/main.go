package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zawazawa0316/designDocConverter/internal/converter"
	"github.com/zawazawa0316/designDocConverter/internal/parser"
	"github.com/zawazawa0316/designDocConverter/internal/writer"
)

func main() {
	// Check if the command line argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: designDocConverter <path_to_excel_file_or_directory>")
		os.Exit(1)
	}
	inputPath := os.Args[1]

	// Check if the path is a directory or a file
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		fmt.Printf("Error accessing input path: %v\n", err)
		os.Exit(1)
	}

	// Process a single file or all .xlsx files in the directory
	if fileInfo.IsDir() {
		// Process all .xlsx files in the directory
		filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error accessing path %s: %v\n", path, err)
				return err
			}
			if !info.IsDir() && strings.HasSuffix(path, ".xlsx") {
				processFile(path)
			}
			return nil
		})
	} else {
		// Process a single file
		processFile(inputPath)
	}
}

// processFile processes a single .xlsx file and converts it to Markdown
func processFile(filePath string) {
	excelData, err := parser.ParseExcelFile(filePath)
	if err != nil {
		fmt.Printf("Failed to parse Excel file %s: %v\n", filePath, err)
		return
	}

	markdownData := converter.ConvertToMarkdown(excelData)
	err = writer.WriteMarkdown(filePath, markdownData)
	if err != nil {
		fmt.Printf("Failed to write Markdown file for %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Markdown file created successfully for %s\n", filePath)
}
