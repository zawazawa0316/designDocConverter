package converter

import (
	"errors"
	"fmt"
	"strings"

	"github.com/zawazawa0316/designDocConverter/internal/parser"
)

// ConvertToMarkdown converts parsed Excel data to Markdown format.
func ConvertToMarkdown(data *parser.ExcelData) (string, error) {
	if data == nil {
		return "", errors.New("no Excel data provided")
	}

	if len(data.Sheets) == 0 {
		return "", errors.New("no sheets found in the Excel data")
	}

	var markdownBuilder strings.Builder

	for _, sheet := range data.Sheets {
		// Convert sheet name to Markdown header
		markdownBuilder.WriteString(fmt.Sprintf("## %s\n\n", sheet.Name))

		// Check if the sheet is empty
		if len(sheet.Table) == 0 {
			continue // Proceed to the next sheet
		}

		// Convert sheet data to Markdown table
		for _, row := range sheet.Table {
			markdownBuilder.WriteString("|")
			for _, cell := range row {
				// Add table data cells
				markdownBuilder.WriteString(fmt.Sprintf(" %s |", cell))
			}

			// End of the row
			markdownBuilder.WriteString("\n")
		}

		// Add an extra newline after each sheet
		markdownBuilder.WriteString("\n")
	}

	return markdownBuilder.String(), nil
}
