# DesignDocConverter (ddc)

## Overview
DesignDocConverter (ddc) is a CLI tool that converts Excel files into Markdown format. It can process a specific xlsx file or all xlsx files within a directory.

## Usage
After building the tool, execute `ddc` using the following command:

```sh
./ddc <path_to_excel_file_or_directory>
```

- `<path_to_excel_file_or_directory>` should be the path to the Excel file you want to convert or a directory containing Excel files.
- The output is generated in the same directory as the input file.
- Sheets are converted to Markdown headings, and cell contents are formatted in a table structure.

### Output Sample
For a simple Excel file with multiple sheets, the Markdown output will look like this:

```
## Sheet1

| A1 Content | B1 Content |
| A2 Content | B2 Content |

## Sheet2

| A1 Content |
| A2 Content |
```

This represents two sheets with their respective cell contents arranged in table format, without header rows.

## Build Instructions
Ensure your environment is set up with Go, then build the project using the following command:

```sh
go build -o ddc github.com/zawazawa0316/designDocConverter/cmd/ddc
```

This will create the `ddc` binary in your project's root directory.

## Running Tests
To run the project tests, use the following command:

```sh
go test github.com/zawazawa0316/designDocConverter/cmd/ddc
```

This command runs tests to verify the accuracy of the conversion process.

---