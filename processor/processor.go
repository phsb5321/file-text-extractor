package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"textractor/config"
	"textractor/filehandler"
)

var fileIndex int

func ProcessDirectory(cfg *config.Config) error {
	inputDir := cfg.InputDir
	outputFile := cfg.OutputFile
	ignoredExts := cfg.IgnoredExts
	onlyExts := cfg.IncludedExts
	maxWordsPerFile := cfg.MaxWordsPerFile
	fileIndex = 0

	// Create initial output file
	outputFile = createNewOutputFile(outputFile, ".txt")

	return filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			err := processFile(
				path,
				ignoredExts,
				onlyExts,
				maxWordsPerFile,
				&outputFile,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func processFile(path string, ignoredExts []string, onlyExts []string, maxWordsPerFile int, outputFile *string) error {
	fileExt := filepath.Ext(path)

	if isFileIgnored(fileExt, ignoredExts) ||
		!isFileIncluded(fileExt, onlyExts) {
		return nil
	}

	content, err := filehandler.ReadFileContent(path)
	if err != nil {
		return err
	}

	if shouldCreateNewFile(content, maxWordsPerFile, *outputFile, fileExt) {
		fileIndex++
		*outputFile = createNewOutputFile(*outputFile, fileExt)
	}

	return appendContentToFiles(content, maxWordsPerFile, outputFile)
}

func isFileIgnored(fileExt string, ignoredExts []string) bool {
	return filehandler.IsIgnoredExtension(fileExt, ignoredExts)
}

func isFileIncluded(fileExt string, onlyExts []string) bool {
	return len(onlyExts) == 0 || filehandler.Contains(onlyExts, fileExt)
}

func shouldCreateNewFile(content string, maxWordsPerFile int, outputFile, fileExt string) bool {
	wordCountExistingFile := 0
	if contentExisting, err := filehandler.ReadFileContent(outputFile); err == nil {
		wordCountExistingFile = filehandler.CountWords(contentExisting)
	}
	wordCountNewContent := filehandler.CountWords(content)
	return (wordCountExistingFile + wordCountNewContent) > maxWordsPerFile
}

func hasDifferentExtension(outputFile, fileExt string) bool {
	return filepath.Ext(outputFile) != fileExt
}

func createNewOutputFile(outputFile, fileExt string) string {
	fileIndex++
	newOutputFile := fmt.Sprintf("%s_%d%s", outputFile, fileIndex, fileExt)
	_ = filehandler.CreateOutputFile(newOutputFile)
	return newOutputFile
}

func appendContentToFiles(content string, maxWordsPerFile int, outputFile *string) error {
	words := strings.Fields(content)

	for i := 0; i < len(words); i += maxWordsPerFile {
		end := i + maxWordsPerFile

		if end > len(words) {
			end = len(words)
		}

		err := filehandler.AppendToOutputFile(*outputFile, strings.Join(words[i:end], " ")) // Modified this line
		if err != nil {
			return err
		}

		if end < len(words) && len(words[end:end+maxWordsPerFile]) > maxWordsPerFile {
			fileIndex++
			*outputFile = createNewOutputFile(*outputFile, filepath.Ext(*outputFile))
		}
	}

	return nil
}
