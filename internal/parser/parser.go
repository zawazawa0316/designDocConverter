package parser

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

// SheetData represents the data of a single sheet.
type SheetData struct {
	Name  string     // Name of the sheet
	Table [][]string // Table data in the sheet
}

// ExcelData represents the entire data of an Excel file.
type ExcelData struct {
	Sheets []SheetData
}

// isEmptyRow checks if a row is empty.
func isEmptyRow(row []string) bool {
	for _, cell := range row {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}

// isColumnEmpty checks if a column is entirely empty.
func isColumnEmpty(rows [][]string, columnIndex int) bool {
	for _, row := range rows {
		if columnIndex < len(row) && strings.TrimSpace(row[columnIndex]) != "" {
			return false
		}
	}
	return true
}

// removeEmptyColumns removes columns that are completely empty.
func removeEmptyColumns(rows [][]string) [][]string {
	var nonEmptyColumnIndexes []int
	if len(rows) == 0 {
		return rows
	}

	// Find indexes of columns that are not empty
	for i := 0; i < len(rows[0]); i++ {
		if !isColumnEmpty(rows, i) {
			nonEmptyColumnIndexes = append(nonEmptyColumnIndexes, i)
		}
	}

	// Keep only data from non-empty columns
	var newRows [][]string
	for _, row := range rows {
		var newRow []string
		for _, idx := range nonEmptyColumnIndexes {
			if idx < len(row) {
				newRow = append(newRow, row[idx])
			} else {
				newRow = append(newRow, "")
			}
		}
		newRows = append(newRows, newRow)
	}
	return newRows
}

// ParseExcelFile parses an Excel file and extracts the data.
func ParseExcelFile(filePath string) (*ExcelData, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %v", err)
	}

	var data ExcelData

	for _, sheetName := range f.GetSheetMap() {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return nil, fmt.Errorf("failed to read sheet '%s': %v", sheetName, err)
		}

		var sheetData SheetData
		sheetData.Name = sheetName
		sheetData.Table = make([][]string, 0)

		for _, row := range rows {
			// Ignore empty rows
			if isEmptyRow(row) {
				continue
			}
			sheetData.Table = append(sheetData.Table, row)
		}

		// Remove empty columns
		sheetData.Table = removeEmptyColumns(sheetData.Table)
		data.Sheets = append(data.Sheets, sheetData)
	}

	return &data, nil
}
