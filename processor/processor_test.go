package processor

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"textractor/config"
)

func TestAll(t *testing.T) {
	t.Run("TestProcessDirectory", TestProcessDirectory)
	t.Run("TestProcessDirectory_NoFiles", TestProcessDirectory_NoFiles)
	t.Run("TestProcessDirectory_IgnoreExtensions", TestProcessDirectory_IgnoreExtensions)
	t.Run("TestProcessDirectory_OnlyIncludeExtensions", TestProcessDirectory_OnlyIncludeExtensions)
	t.Run("TestProcessDirectory_WordCountExceedsMax", TestProcessDirectory_WordCountExceedsMax)
}

// TestProcessDirectory tests the core function of the processor package.
// It creates a directory structure with multiple files of different types,
// some of which should be processed based on their extensions and others which should be ignored.
// It then calls ProcessDirectory with this directory structure and checks that the output file contains the expected contents.
func TestProcessDirectory(t *testing.T) {
	// This test needs a real directory and files to work with, so set up some test data.
	err := os.Mkdir("test_dir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	// Create a dummy text file in the directory.
	err = os.WriteFile("test_dir/test.txt", []byte("Test data."), 0644)
	if err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		InputDir:        "test_dir",
		OutputFile:      "output.txt",
		MaxWordsPerFile: 10,
		IgnoredExts:     []string{".pdf", ".docx"},
		IncludedExts:    []string{".txt"},
	}

	// Now we can call ProcessDirectory with our test directory.
	err = ProcessDirectory(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the output file exists.
	_, err = os.Stat("output.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Check that the output.txt file contains the expected contents.
	expectedContent := "Test data."
	content, err := ioutil.ReadFile("output.txt")
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != expectedContent {
		t.Errorf("Output file content mismatch. Expected: %s, Got: %s", expectedContent, string(content))
	}

	// Check that only the output.txt file was created.
	files, err := ioutil.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	numOutputFiles := 0
	for _, file := range files {
		if file.Name() == "output.txt" {
			numOutputFiles++
		}
	}

	if numOutputFiles != 1 {
		t.Errorf("Unexpected number of output files. Expected: 1, Got: %d", numOutputFiles)
	}

	// Clean up test files.

	// Remove the test directory.
	err = os.RemoveAll("test_dir")
	if err != nil {
		t.Fatal(err)
	}

	// Remove the output file.
	err = os.Remove("output.txt")
	if err != nil {
		t.Fatal(err)
	}
}

// TestProcessDirectory_NoFiles tests the case where the input directory has no files.
// It checks that no output file is created.
func TestProcessDirectory_NoFiles(t *testing.T) {
	err := os.Mkdir("test_dir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		InputDir:        "test_dir",
		OutputFile:      "output.txt",
		MaxWordsPerFile: 10,
		IgnoredExts:     []string{".pdf", ".docx"},
		IncludedExts:    []string{".txt"},
	}

	err = ProcessDirectory(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Check that no output file is created.
	_, err = os.Stat("output.txt")
	if !os.IsNotExist(err) {
		t.Errorf("Output file should not exist, but found: %v", err)
	}

	// Clean up test directory.
	err = os.RemoveAll("test_dir")
	if err != nil {
		t.Fatal(err)
	}
}

// TestProcessDirectory_IgnoreExtensions tests the case where files with ignored extensions are present in the input directory.
// It checks that files with ignored extensions are not processed.
func TestProcessDirectory_IgnoreExtensions(t *testing.T) {
	err := os.Mkdir("test_dir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("test_dir/test.txt", []byte("Test data."), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("test_dir/test.pdf", []byte("PDF data."), 0644)
	if err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		InputDir:        "test_dir",
		OutputFile:      "output.txt",
		MaxWordsPerFile: 10,
		IgnoredExts:     []string{".pdf", ".docx"},
		IncludedExts:    []string{".txt"},
	}

	err = ProcessDirectory(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the output file exists.
	_, err = os.Stat("output.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Check that the output.txt file does not contain the ignored PDF data.
	content, err := ioutil.ReadFile("output.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "Test data."
	if string(content) != expectedContent {
		t.Errorf("Output file content mismatch. Expected: %s, Got: %s", expectedContent, string(content))
	}

	// Clean up test files.
	err = os.RemoveAll("test_dir")
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove("output.txt")
	if err != nil {
		t.Fatal(err)
	}
}

// TestProcessDirectory_OnlyIncludeExtensions tests the case where only files with specific extensions are included for processing.
// It checks that only files with the specified extensions are processed.
func TestProcessDirectory_OnlyIncludeExtensions(t *testing.T) {
	err := os.Mkdir("test_dir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("test_dir/test.txt", []byte("Test data."), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("test_dir/test.pdf", []byte("PDF data."), 0644)
	if err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		InputDir:        "test_dir",
		OutputFile:      "output.txt",
		MaxWordsPerFile: 10,
		IgnoredExts:     []string{".pdf", ".docx"},
		IncludedExts:    []string{".txt"},
	}

	err = ProcessDirectory(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the output file exists.
	_, err = os.Stat("output.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Check that the output.txt file does not contain the ignored PDF data.
	content, err := ioutil.ReadFile("output.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "Test data."
	if string(content) != expectedContent {
		t.Errorf("Output file content mismatch. Expected: %s, Got: %s", expectedContent, string(content))
	}

	// Clean up test files.
	err = os.RemoveAll("test_dir")
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove("output.txt")
	if err != nil {
		t.Fatal(err)
	}
}

// TestProcessDirectory_WordCountExceedsMax tests the case where the word count of a file exceeds the maximum word count per file.
// It checks that a new output file is created to accommodate the content of the file.
func TestProcessDirectory_WordCountExceedsMax(t *testing.T) {
	// Clean up test directory if it already exists
	_ = os.RemoveAll("test_dir")

	err := os.Mkdir("test_dir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	// Create a large text file with more than the maximum word count.
	content := "This is a large text file with more words than the maximum word count per file."
	err = os.WriteFile("test_dir/large.txt", []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		InputDir:        "test_dir",
		OutputFile:      "output.txt",
		MaxWordsPerFile: 10,
		IgnoredExts:     []string{},
		IncludedExts:    []string{".txt"},
	}

	err = ProcessDirectory(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the output files exist.
	numOutputFiles := len(cfg.IncludedExts)
	for i := 0; i < numOutputFiles; i++ {
		expectedOutputFile := fmt.Sprintf("output%s.txt", getOutputFileIndex(i))
		_, err := os.Stat(expectedOutputFile)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Check the content of each output file.
	expectedContent := "This is a large text file with more words than the maximum word count per file."
	chunkSize := len(expectedContent) / numOutputFiles

	for i := 0; i < numOutputFiles; i++ {
		expectedOutputFile := fmt.Sprintf("output%s.txt", getOutputFileIndex(i))
		content, err := ioutil.ReadFile(expectedOutputFile)
		if err != nil {
			t.Fatal(err)
		}

		expectedChunk := expectedContent[i*chunkSize : (i+1)*chunkSize]
		if string(content) != expectedChunk {
			t.Errorf("Output file content mismatch. Expected: %s, Got: %s", expectedChunk, string(content))
		}
	}

	// Clean up test files.
	err = os.RemoveAll("test_dir")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < numOutputFiles; i++ {
		expectedOutputFile := fmt.Sprintf("output%s.txt", getOutputFileIndex(i))
		err = os.Remove(expectedOutputFile)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// getOutputFileIndex returns the index string for the output files based on the given number.
func getOutputFileIndex(num int) string {
	if num == 0 {
		return ""
	}
	return fmt.Sprintf("_%d", num)
}
