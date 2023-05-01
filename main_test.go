package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestProcessDirectory(t *testing.T) {
	// Create temporary directory and file for testing
	tempDir, err := ioutil.TempDir("", "test_data_*")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempFile, err := ioutil.TempFile(tempDir, "test1_*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	tempFile.Close()

	outputFile, err := ioutil.TempFile("", "test_output_*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary output file: %v", err)
	}
	outputFileName := outputFile.Name()
	defer os.Remove(outputFileName)

	err = processDirectory(tempDir, outputFileName, []string{})
	if err != nil {
		t.Fatalf("Error processing directory: %v", err)
	}

	// Check if output file was created
	if _, err := os.Stat(outputFileName); os.IsNotExist(err) {
		t.Fatalf("Output file not created: %v", err)
	}

	content, err := ioutil.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	expectedSubstring := fmt.Sprintf("File: %s", filepath.Base(tempFile.Name()))
	if !strings.Contains(string(content), expectedSubstring) {
		t.Fatalf("Output file content does not contain expected substring: %s", expectedSubstring)
	}
}

func TestProcessDirectoryWithIgnoredExtensions(t *testing.T) {
	// Create temporary directory and file for testing
	tempDir, err := ioutil.TempDir("", "test_data_*")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempFile, err := ioutil.TempFile(tempDir, "test1_*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	tempFile.Close()

	outputFile, err := ioutil.TempFile("", "test_output_*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary output file: %v", err)
	}
	outputFileName := outputFile.Name()
	defer os.Remove(outputFileName)

	err = processDirectory(tempDir, outputFileName, []string{".txt"})
	if err != nil {
		t.Fatalf("Error processing directory: %v", err)
	}

	// Check if output file was created
	if _, err := os.Stat(outputFileName); os.IsNotExist(err) {
		t.Fatalf("Output file not created: %v", err)
	}

	content, err := ioutil.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	expectedSubstring := fmt.Sprintf("File: %s", filepath.Base(tempFile.Name()))
	if strings.Contains(string(content), expectedSubstring) {
		t.Fatalf("Output file content should not contain ignored file substring: %s", expectedSubstring)
	}
}

func TestProcessFile(t *testing.T) {
	// Create temporary directory and file for testing
	tempDir, err := ioutil.TempDir("", "test_data_*")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempFile, err := ioutil.TempFile(tempDir, "test1_*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	tempFileContent := "This is a test file."
	tempFile.WriteString(tempFileContent)
	tempFile.Close()

	outputFileName := filepath.Join(tempDir, "test_output.txt")

	fileInfo, err := os.Stat(tempFile.Name())
	if err != nil {
		t.Fatalf("Error stating input file: %v", err)
	}

	err = processFile(tempFile.Name(), fileInfo, nil, outputFileName, []string{})
	if err != nil {
		t.Fatalf("Error processing file: %v", err)
	}

	// Check if output file was created
	if _, err := os.Stat(outputFileName); os.IsNotExist(err) {
		t.Fatalf("Output file not created: %v", err)
	}

	content, err := ioutil.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	expectedSubstring := tempFileContent
	if !strings.Contains(string(content), expectedSubstring) {
		t.Fatalf("Output file content does not contain expected substring: %s", expectedSubstring)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running tests...")

	exitVal := m.Run()

	fmt.Println("Cleaning up...")

	os.Exit(exitVal)
}
