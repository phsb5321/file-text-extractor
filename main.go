package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var inputDir string
	flag.StringVar(&inputDir, "d", "", "Directory to process")

	var outputFileName string
	flag.StringVar(&outputFileName, "o", "", "Output file name")

	var ignoredExtensions string
	flag.StringVar(&ignoredExtensions, "i", "", "Comma separated list of extensions to ignore")

	flag.Parse()

	ignoredExtList := strings.Split(ignoredExtensions, ",")
	// Print ignored extensions
	fmt.Printf("Ignored extensions: %v\n", ignoredExtList)

	if inputDir == "" {
		fmt.Println("Input directory is required")
		os.Exit(1)
	}

	if outputFileName == "" {
		fmt.Println("Output file name is required")
		os.Exit(1)
	}

	err := processDirectory(inputDir, outputFileName, ignoredExtList)
	if err != nil {
		fmt.Printf("Error processing directory: %v\n", err)
		os.Exit(1)
	}
}

func processDirectory(inputDir string, outputFileName string, ignoredExtList []string) error {
	return filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		return processFile(path, info, err, outputFileName, ignoredExtList)
	})
}

func processFile(path string, info os.FileInfo, err error, outputFileName string, ignoredExtList []string) error {
	if err != nil {
		fmt.Printf("Error processing file %s: %v\n", path, err)
		return nil
	}

	if !info.IsDir() {
		ext := filepath.Ext(path)

		// Add "." to ignored extensions if it is missing
		for i, ignoredExt := range ignoredExtList {
			if !strings.HasPrefix(ignoredExt, ".") {
				ignoredExtList[i] = "." + ignoredExt
			}
		}

		for _, ignoredExt := range ignoredExtList {
			if ignoredExt == ext {
				return nil
			}
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			return nil
		}

		outFile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening output file %s: %v\n", outputFileName, err)
			return nil
		}
		defer outFile.Close()

		header := fmt.Sprintf("\n---\nFile: %s\nDirectory: %s\n---\n", info.Name(), filepath.Dir(path))
		if _, err := outFile.WriteString(header); err != nil {
			fmt.Printf("Error writing header to output file: %v\n", err)
			return nil
		}

		if _, err := outFile.Write(content); err != nil {
			fmt.Printf("Error writing content to output file: %v\n", err)
			return nil
		}
	}

	return nil
}
