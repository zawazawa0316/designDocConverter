package converter

import (
	"testing"

	"github.com/zawazawa0316/designDocConverter/internal/parser"
)

// TestConvertToMarkdownWithValidData tests conversion of valid Excel data to Markdown.
func TestConvertToMarkdownWithValidData(t *testing.T) {
	data := &parser.ExcelData{
		Sheets: []parser.SheetData{
			{
				Name: "Sheet1",
				Table: [][]string{
					{"Data1", "Data2"},
					{"Data3", "Data4"},
				},
			},
		},
	}

	expectedMarkdown := "## Sheet1\n\n" +
		"| Data1 | Data2 |\n" +
		"| Data3 | Data4 |\n\n"

	result, err := ConvertToMarkdown(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expectedMarkdown {
		t.Errorf("unexpected result: got %v, want %v", result, expectedMarkdown)
	}
}

// TestConvertToMarkdownWithEmptyData tests conversion with empty Excel data.
func TestConvertToMarkdownWithEmptyData(t *testing.T) {
	data := &parser.ExcelData{}

	_, err := ConvertToMarkdown(data)
	if err == nil {
		t.Fatalf("expected an error, but got none")
	}
}

// TestConvertToMarkdownWithNilData tests conversion with nil Excel data.
func TestConvertToMarkdownWithNilData(t *testing.T) {
	_, err := ConvertToMarkdown(nil)
	if err == nil {
		t.Fatalf("expected an error, but got none")
	}
}
