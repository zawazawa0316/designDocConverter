package writer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TestWriteMarkdownWithValidData tests writing valid markdown data to a file.
func TestWriteMarkdownWithValidData(t *testing.T) {
	markdownData := "This is a test.\n\n" +
		"| Cell1 | Cell2 |\n" +
		"| Cell3 | Cell4 |"
	tempDir, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(tempDir) // Clean up

	outputPath := filepath.Join(tempDir, "test.md")
	err := WriteMarkdown(outputPath, markdownData)
	if err != nil {
		t.Fatalf("Failed to write markdown: %v", err)
	}

	// Check if the file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("File not created: %s, error: %v", outputPath, err)
	}

	// Read the file and compare its contents
	writtenData, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read the written file: %v", err)
	}
	if string(writtenData) != markdownData {
		t.Errorf("Written data does not match input markdown data")
	}
}

// TestWriteMarkdownWithInvalidPath tests handling of invalid file paths.
func TestWriteMarkdownWithInvalidPath(t *testing.T) {
	markdownData := "# Test Markdown"
	invalidPath := filepath.Join("invalid", "path", "test.md")
	err := WriteMarkdown(invalidPath, markdownData)
	if err == nil {
		t.Fatalf("Expected an error for an invalid path, but did not get one")
	}
}

// TestWriteMarkdownWithEmptyData tests writing an empty markdown string.
func TestWriteMarkdownWithEmptyData(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(tempDir)

	outputPath := filepath.Join(tempDir, "empty.md")
	err := WriteMarkdown(outputPath, "")
	if err != nil {
		t.Fatalf("Failed to write empty markdown: %v", err)
	}

	// Check if the file is created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("File not created: %s", outputPath)
	}
}
