package ddc

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// copyFile copies a file from the source path to the destination path.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func TestCLIRunWithInsufficientArgs(t *testing.T) {
	cli := CLI{
		Stdout: new(bytes.Buffer),
		Stderr: new(bytes.Buffer),
	}

	args := []string{"ddc"}
	err := cli.Run(args)

	if err == nil {
		t.Errorf("Expected an error for insufficient arguments, but got none")
	}
}

// TestCLIRunWithSingleFile tests the CLI's ability to process a single Excel file.
func TestCLIRunWithSingleFile(t *testing.T) {
	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "ddc_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy testdata.xlsx to the temporary directory
	testFilePath := filepath.Join(tempDir, "testdata.xlsx")
	originalTestFilePath := "../../testdata/valid.xlsx"
	if err := copyFile(originalTestFilePath, testFilePath); err != nil {
		t.Fatalf("Failed to copy testdata.xlsx to temp directory: %v", err)
	}

	cli := CLI{
		Stdout: new(bytes.Buffer),
		Stderr: new(bytes.Buffer),
	}

	args := []string{"ddc", testFilePath}
	err = cli.Run(args)

	if err != nil {
		t.Errorf("Failed to process a valid file: %v", err)
	}

	// TODO
	// Verify the existence and content of the output Markdown file
}

// TestCLIRunWithDirectory tests the CLI's ability to process all Excel files in a directory.
func TestCLIRunWithDirectory(t *testing.T) {
	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "ddc_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy testdata.xlsx to the temporary directory as testdata1.xlsx and testdata2.xlsx
	originalTestFilePath := "../../testdata/valid.xlsx"
	testFilePath1 := filepath.Join(tempDir, "testdata1.xlsx")
	testFilePath2 := filepath.Join(tempDir, "testdata2.xlsx")
	if err := copyFile(originalTestFilePath, testFilePath1); err != nil {
		t.Fatalf("Failed to copy testdata.xlsx to temp directory as testdata1.xlsx: %v", err)
	}
	if err := copyFile(originalTestFilePath, testFilePath2); err != nil {
		t.Fatalf("Failed to copy testdata.xlsx to temp directory as testdata2.xlsx: %v", err)
	}

	cli := CLI{
		Stdout: new(bytes.Buffer),
		Stderr: new(bytes.Buffer),
	}

	args := []string{"ddc", tempDir}
	err = cli.Run(args)

	if err != nil {
		t.Errorf("Failed to process files in a valid directory: %v", err)
	}

	// TODO
	// Verify the existence and content of the output Markdown files for both testdata1.xlsx and testdata2.xlsx
}

func TestCLIRunWithInvalidPath(t *testing.T) {
	cli := CLI{
		Stdout: new(bytes.Buffer),
		Stderr: new(bytes.Buffer),
	}

	args := []string{"ddc", "/invalid/path"}
	err := cli.Run(args)

	if err == nil {
		t.Errorf("Expected an error for an invalid path, but got none")
	}
}
