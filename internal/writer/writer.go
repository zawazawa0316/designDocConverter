package writer

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteMarkdown writes markdown data to a .md file in the same directory as the input file.
func WriteMarkdown(inputFilePath, markdownData string) error {
	// Determine the output file path
	outputFilePath := filepath.Join(filepath.Dir(inputFilePath), filepath.Base(inputFilePath)+".md")

	// Check if the output file path is valid
	if outputFilePath == "" {
		return fmt.Errorf("invalid output file path derived from input file: %s", inputFilePath)
	}

	// Write the markdown data to the file
	err := os.WriteFile(outputFilePath, []byte(markdownData), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to %s: %v", outputFilePath, err)
	}

	return nil
}
