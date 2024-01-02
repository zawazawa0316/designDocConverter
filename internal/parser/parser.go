package parser

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// ExcelParser handles the parsing of Excel files.
type ExcelParser struct {
	// Any configuration or dependencies needed can be added here.
}

// NewExcelParser creates a new instance of ExcelParser.
func NewExcelParser() *ExcelParser {
	return &ExcelParser{}
}

// ListFiles returns a list of Excel files in the specified path.
func (p *ExcelParser) ListFiles(inputPath string) ([]string, error) {
	var files []string

	// Check if inputPath is a file or folder
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error stating input path: %w", err)
	}

	if fileInfo.IsDir() {
		// Parse all Excel files in the folder
		err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".xlsx") {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("error walking through the folder: %w", err)
		}
	} else {
		// Check if the specified file is an Excel file
		if !strings.HasSuffix(fileInfo.Name(), ".xlsx") {
			return nil, errors.New("specified file is not an Excel file")
		}
		files = append(files, inputPath)
	}

	if len(files) == 0 {
		return nil, errors.New("no Excel files found in the specified path")
	}

	return files, nil
}

// ParseExcelFile reads the content of an Excel file and returns the sheet data.
func (p *ExcelParser) ParseExcelFile(filePath string) (map[string][][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening Excel file: %w", err)
	}

	sheetData := make(map[string][][]string)

	// Read all sheets in the Excel file
	sheetList := f.GetSheetList()
	for _, sheetName := range sheetList {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, fmt.Errorf("error reading sheet %s: %w", sheetName, err)
		}
		sheetData[sheetName] = rows
	}

	return sheetData, nil
}
