package ddc

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
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

	// Read the expected output from valid.md
	expectedOutput, err := ioutil.ReadFile("../../testdata/valid.md")
	if err != nil {
		t.Fatalf("Failed to read the expected output file: %v", err)
	}

	// Read the actual output generated by the CLI
	outputPath := filepath.Join(tempDir, "testdata.md") // Adjust the output file name as per CLI implementation
	actualOutput, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read the actual output file: %v", err)
	}

	// Compare the actual output with the expected output and display both if they don't match
	if !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Errorf("The actual output does not match the expected output\nActual:\n%s\nExpected:\n%s",
			string(actualOutput), string(expectedOutput))
	}

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

	// Read the expected output from valid.md
	expectedOutput, err := ioutil.ReadFile("../../testdata/valid.md")
	if err != nil {
		t.Fatalf("Failed to read the expected output file: %v", err)
	}

	// Define a helper function to verify the output
	verifyOutput := func(testFileName string, expectedOutput []byte) {
		outputPath := filepath.Join(tempDir, testFileName) // Adjust as per CLI implementation
		actualOutput, err := ioutil.ReadFile(outputPath)
		if err != nil {
			t.Fatalf("Failed to read the actual output file for %s: %v", testFileName, err)
		}
		if !reflect.DeepEqual(actualOutput, expectedOutput) {
			t.Errorf("The actual output for %s does not match the expected output\nActual:\n%s\nExpected:\n%s",
				testFileName, string(actualOutput), string(expectedOutput))
		}
	}

	// Verify the output for testdata1.xlsx and testdata2.xlsx
	verifyOutput("testdata1.md", expectedOutput)
	verifyOutput("testdata2.md", expectedOutput)
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
