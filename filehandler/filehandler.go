package filehandler

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// WriteContentToFile writes the given content to the specified file.
func WriteContentToFile(fileName, content string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// ReadFileContent reads and returns the content of the file at the specified path.
func ReadFileContent(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// IsIgnoredExtension checks if the given file extension is present in the ignored extensions list.
func IsIgnoredExtension(fileExt string, ignoredExts []string) bool {
	for _, ignoredExt := range ignoredExts {
		if strings.EqualFold(ignoredExt, fileExt) {
			return true
		}
	}
	return false
}

// AppendDotToExtensions adds a "." prefix to the file extensions that don't have it already.
func AppendDotToExtensions(extList []string) []string {
	for i, ext := range extList {
		if !strings.HasPrefix(ext, ".") {
			extList[i] = "." + ext
		}
	}
	return extList
}

// Contains checks if the given slice contains the given item.
func Contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

// AppendToOutputFile appends the given content to the specified output file.
func AppendToOutputFile(outputFile, content string) error {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// CreateOutputFile creates the output file with the given path if it doesn't exist.
func CreateOutputFile(outputFile string) error {
	_, err := os.Stat(outputFile)
	if os.IsNotExist(err) {
		dir := filepath.Dir(outputFile)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = os.Create(outputFile)
		if err != nil {
			return err
		}
	}
	return nil
}

// Exists checks if the file with the given path exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// CountWords counts the number of words in the given content.
func CountWords(content string) int {
	wordCount := 0
	inWord := false
	for _, r := range content {
		if isWordChar(r) {
			if !inWord {
				inWord = true
				wordCount++
			}
		} else {
			inWord = false
		}
	}
	return wordCount
}

// isWordChar returns true if the given rune is a valid word character.
func isWordChar(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

// FileExists checks if the file with the given path exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
