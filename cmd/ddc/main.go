package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/your-username/ddc/internal/converter"
	"github.com/your-username/ddc/internal/output"
	"github.com/your-username/ddc/internal/parser"
)

func main() {
	// Define command-line flags
	inputPath := flag.String("i", "", "Path to the input Excel file or folder")
	outputPath := flag.String("o", "", "Path to the output Markdown file or folder")
	format := flag.String("f", "md", "Output format (md, html, etc.)")
	configPath := flag.String("c", "", "Path to the configuration file")

	// Parse command-line flags
	flag.Parse()

	// Validate required flags
	if *inputPath == "" || *outputPath == "" {
		fmt.Println("Error: Both input and output paths are required.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Create parser
	excelParser := parser.NewExcelParser()

	// Parse Excel file or folder
	files, err := excelParser.Parse(*inputPath)
	if err != nil {
		log.Fatal("Error parsing Excel file/folder:", err)
	}

	// Create converter
	mdConverter := converter.NewMarkdownConverter()

	// Convert Excel data to Markdown
	mdContent, err := mdConverter.Convert(files)
	if err != nil {
		log.Fatal("Error converting Excel data to Markdown:", err)
	}

	// Create output handler
	outputHandler := output.NewOutputHandler()

	// Output Markdown content to file or folder
	err = outputHandler.Output(mdContent, *outputPath, *format)
	if err != nil {
		log.Fatal("Error writing Markdown content to file/folder:", err)
	}

	fmt.Println("Conversion completed successfully!")
}
