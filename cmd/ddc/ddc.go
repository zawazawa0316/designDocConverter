package ddc

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/zawazawa0316/designDocConverter/internal/converter"
	"github.com/zawazawa0316/designDocConverter/internal/parser"
	"github.com/zawazawa0316/designDocConverter/internal/writer"
)

type CLI struct {
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
}

func (cli *CLI) Run(args []string) error {
	if len(args) < 2 {
		fmt.Fprintln(cli.Stderr, "Usage: designDocConverter <path_to_excel_file_or_directory>")
		return fmt.Errorf("insufficient arguments")
	}
	inputPath := args[1]

	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		fmt.Fprintf(cli.Stderr, "Error accessing input path: %v\n", err)
		return err
	}

	if fileInfo.IsDir() {
		return filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Fprintf(cli.Stderr, "Error accessing path %s: %v\n", path, err)
				return err
			}
			if !info.IsDir() && strings.HasSuffix(path, ".xlsx") {
				return processFile(cli, path)
			}
			return nil
		})
	} else {
		return processFile(cli, inputPath)
	}
}

func processFile(cli *CLI, filePath string) error {
	excelData, err := parser.ParseExcelFile(filePath)
	if err != nil {
		fmt.Fprintf(cli.Stderr, "Failed to parse Excel file %s: %v\n", filePath, err)
		return err
	}

	markdownData, err := converter.ConvertToMarkdown(excelData)
	if err != nil {
		fmt.Fprintf(cli.Stderr, "Failed to convert Excel file to Markdown %s: %v\n", filePath, err)
		return err
	}

	err = writer.WriteMarkdown(filePath, markdownData)
	if err != nil {
		fmt.Fprintf(cli.Stderr, "Failed to write Markdown file for %s: %v\n", filePath, err)
		return err
	}

	fmt.Fprintf(cli.Stdout, "Markdown file created successfully for %s\n", filePath)
	return nil
}
