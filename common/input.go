// Package common provides utility functions shared across the project.
package common

import (
	"bufio"
	"os"

	"gopkg.in/yaml.v2"
)

// Config
type Config struct {
	InputFile string `yaml:"inputFile"`
	LogLevel  string `yaml:"logLevel"`
}

// readConfig reads the YAML configuration file and returns the config
func ReadConfig(paths ...string) (Config, error) {
	var cfg Config
	configFilePath := "config.yaml" // Default configuration file

	if len(paths) > 0 {
		configFilePath = paths[0]
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// ReadInputFile reads contents of a file and returns them as a string.
// The function reads from `input.txt` by default unless the env var
// "ADVENT_OF_CODE_TEST" has been set to "TRUE".
func ReadInputFile(cfg Config) (string, error) {
	data, err := os.ReadFile(cfg.InputFile)

	// Go error handling.
	if err != nil {
		return "", err
	}

	// Convert data from byte slice to a string and return.
	return string(data), nil
}

// ReadInputFileAs2DSlice reads contents of a file and returns them as a 2D
// slice of runes. Useful for taking input as a 2D grid with coordinates.
// Will remove empty lines.
func ReadInputFileAs2DSlice(cfg Config) ([][]rune, error) {
	file, err := os.Open(cfg.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Check if line is not empty
			lines = append(lines, []rune(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
