## Overview

`ddc` (designDocConverter) is a command-line tool designed to convert Excel files to Markdown format. With this tool, you can easily transform Excel-based design documents into Markdown format.

## Installation

1. Download the appropriate binary for your platform from the GitHub release page.
2. Place the downloaded binary in a suitable location and add it to your executable path.

## Usage

```bash
ddc [OPTIONS] input_path output_path
```
input_path: The path to the Excel file or folder to be converted.
output_path: The path to the Markdown file or folder where the conversion result will be stored.
Options
-h, --help: Display the help message.

-f, --format FORMAT: Specify the output format. The default is Markdown, but other formats are also available.
```bash
ddc -f html input.xlsx output.html
```

-c, --config FILE: Specify a configuration file. You can use the config file to define conversion options.
```bash
ddc -c config.json input.xlsx output.md
```
## Examples
1. Convert a single file:

```bash
ddc input.xlsx output.md
```
2. Convert all Excel files in a folder:

```bash
ddc input_folder/ output_folder/
```

## Notes
* Ensure that Excel files are in the correct format.
* In case of errors during conversion, check the error messages for details and troubleshoot accordingly.