package config

import (
	"errors"
	"fmt"
	"math"
	"os"

	"github.com/spf13/pflag"
)

// Config represents the configuration options for the program
type Config struct {
	InputDir        string   // the input directory to search for files
	OutputFile      string   // the name of the output file
	IgnoredExts     []string // a list of file extensions to ignore
	IncludedExts    []string // a list of file extensions to only include
	MaxWordsPerFile int      // the maximum number of words per output file
	IncludedDirs    []string // New
	ExcludedDirs    []string // New
}

const (
	MAX_WORDS_PER_FILE = math.MaxInt64
)

// ParseCommandLineArguments parses the command-line arguments and returns the configuration options
func ParseCommandLineArguments(args []string) (*Config, error) {
	// create a new Config struct
	cfg := &Config{}

	// create a new FlagSet and add flags for InputDir, OutputFile, IgnoredExts, and OnlyExt
	flags := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	flags.StringVarP(&cfg.InputDir, "input-directory", "d", "", "input directory")
	flags.StringVarP(&cfg.OutputFile, "output-file", "o", "output.txt", "output file name")
	flags.StringSliceVarP(&cfg.IgnoredExts, "ignored-exts", "i", []string{".jpg", ".png"}, "comma-separated list of ignored file extensions")
	flags.StringSliceVar(&cfg.IncludedExts, "only", []string{}, "comma-separated list of file extensions to process exclusively")
	flags.IntVarP(&cfg.MaxWordsPerFile, "max-words-per-file", "w", MAX_WORDS_PER_FILE, "maximum number of words per output file")

	// use catchPanic to recover from any panics that might occur while parsing flags
	err := catchPanic(func() {
		if err := flags.Parse(args); err != nil {
			panic(err)
		}
	})

	// check for errors during flag parsing and validation
	if err != nil {
		return nil, fmt.Errorf("failed to parse command-line arguments: %s", err)
	}
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// catchPanic catches any panics that occur during the execution of f and returns them as an error
func catchPanic(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	f()
	return
}

// validateConfig validates the configuration options
func validateConfig(cfg *Config) error {
	// ensure that InputDir is not empty
	if cfg.InputDir == "" {
		return errors.New("input directory is required")
	}

	// check if InputDir exists
	if _, err := os.Stat(cfg.InputDir); os.IsNotExist(err) {
		return fmt.Errorf("input directory does not exist: %s", cfg.InputDir)
	}

	return nil
}
