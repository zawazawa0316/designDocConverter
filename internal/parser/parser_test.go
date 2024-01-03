package parser

import (
	"path/filepath"
	"testing"
)

// TestParseExcelFileWithValidFile tests parsing a valid Excel file.
func TestParseExcelFileWithValidFile(t *testing.T) {
	validFilePath := filepath.Join("..", "..", "testdata", "valid.xlsx")

	data, err := ParseExcelFile(validFilePath)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(data.Sheets) == 0 {
		t.Errorf("Expected non-empty sheets, got empty")
	}
}

// TestParseExcelFileWithNonExistentFile tests parsing a non-existent file.
func TestParseExcelFileWithNonExistentFile(t *testing.T) {
	nonExistentFilePath := filepath.Join("..", "..", "testdata", "nonexistent.xlsx")

	_, err := ParseExcelFile(nonExistentFilePath)

	if err == nil {
		t.Errorf("Expected an error for non-existent file, got none")
	}
}

// TestParseExcelFileWithCorruptedFile tests parsing a corrupted Excel file.
func TestParseExcelFileWithCorruptedFile(t *testing.T) {
	corruptedFilePath := filepath.Join("..", "..", "testdata", "corrupted.xlsx")

	_, err := ParseExcelFile(corruptedFilePath)

	if err == nil {
		t.Errorf("Expected an error for corrupted file, got none")
	}
}
