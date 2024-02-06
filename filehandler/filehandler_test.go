package filehandler

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("TestWriteContentToFile", TestWriteContentToFile)
	t.Run("TestIsIgnoredExtension", TestIsIgnoredExtension)
	t.Run("TestAppendDotToExtensions", TestAppendDotToExtensions)
}

func TestWriteContentToFile(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		content       string
		expectedError error
	}{
		{
			name:     "valid input",
			fileName: "test_output.txt",
			content:  "This is a test content",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := WriteContentToFile(test.fileName, test.content)

			if err != test.expectedError {
				t.Errorf("Unexpected error, expected %v, got %v", test.expectedError, err)
			}

			verifyFileContentAndCleanup(t, test)
		})
	}
}

func verifyFileContentAndCleanup(t *testing.T, test struct {
	name          string
	fileName      string
	content       string
	expectedError error
}) {
	if test.expectedError == nil {
		// Verify the content of the file
		content, err := ioutil.ReadFile(test.fileName)
		if err != nil {
			t.Fatalf("Error reading output file: %v", err)
		}

		if string(content) != test.content {
			t.Errorf("Unexpected content, expected %s, got %s", test.content, string(content))
		}

		// Clean up
		err = os.Remove(test.fileName)
		if err != nil {
			t.Errorf("Error cleaning up test file: %v", err)
		}
	}
}

func TestIsIgnoredExtension(t *testing.T) {
	tests := []struct {
		name          string
		fileExt       string
		ignoredExts   []string
		expectedValue bool
	}{
		{
			name:          "not ignored",
			fileExt:       ".txt",
			ignoredExts:   []string{".png", ".jpg"},
			expectedValue: false,
		},
		{
			name:          "ignored",
			fileExt:       ".png",
			ignoredExts:   []string{".png", ".jpg"},
			expectedValue: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsIgnoredExtension(test.fileExt, test.ignoredExts)

			if result != test.expectedValue {
				t.Errorf("Unexpected value, expected %v, got %v", test.expectedValue, result)
			}
		})
	}
}

func TestAppendDotToExtensions(t *testing.T) {
	tests := []struct {
		name          string
		extList       []string
		expectedValue []string
	}{
		{
			name:          "mixed input",
			extList:       []string{".txt", "jpg", "png"},
			expectedValue: []string{".txt", ".jpg", ".png"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := AppendDotToExtensions(test.extList)

			for i, ext := range result {
				if ext != test.expectedValue[i] {
					t.Errorf("Unexpected value, expected %v, got %v", test.expectedValue[i], ext)
				}
			}
		})
	}
}
