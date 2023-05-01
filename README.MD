# File Text Extractor

## Introduction

File Text Extractor is a CLI tool developed in Go that processes files in a given directory and extracts their text content to an output file. It can also ignore files with certain extensions.

## Installation

1. Install Go on your machine. Instructions can be found [here](https://golang.org/doc/install).
2. Clone the repository to your local machine: `git clone https://github.com/username/repo-name.git`
3. Navigate to the project directory: `cd repo-name`
4. Build the executable: `go build`

## Usage

The tool can be executed with the following command:

`./file-text-extractor -d <directory> -o <output file> -i <ignored extensions>`

Where:

- `-d`: directory to process
- `-o`: output file
- `-i`: comma-separated list of file extensions to ignore

### Example usage

Process all files in the directory `/home/user/documents` and output the results to a file named `output.txt`, ignoring files with extensions `.pdf` and `.docx`:

`./file-text-extractor -d /home/user/documents -o output.txt -i pdf,docx`

## Running tests

Tests can be run with the following command:

`go test -cover`

## Contributing

Contributions to this project are always welcome! Here are some ways to contribute:

- Report a bug by opening an issue
- Fix an issue by opening a pull request
- Add new features by opening a pull request

## License

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
