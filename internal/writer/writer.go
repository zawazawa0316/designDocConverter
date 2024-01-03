package writer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteMarkdown writes markdown data to a .md file in the same directory as the input file.
func WriteMarkdown(inputFilePath, markdownData string) error {
	// Remove the extension from the input file name and add .md extension
	baseName := strings.TrimSuffix(filepath.Base(inputFilePath), filepath.Ext(inputFilePath))
	outputFilePath := filepath.Join(filepath.Dir(inputFilePath), baseName+".md")

	// Write the markdown data to the file
	err := os.WriteFile(outputFilePath, []byte(markdownData), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to %s: %v", outputFilePath, err)
	}

	return nil
}
