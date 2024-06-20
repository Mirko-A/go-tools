package io

import (
	"io"
	"os"
	"strings"
)

func ReadToString(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func ReadLines(filePath string) ([]string, error) {
	s, err := ReadToString(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(s, "\n"), nil
}

func ReadLinesCRLF(filePath string) ([]string, error) {
	s, err := ReadToString(filePath)
	if err != nil {
		return nil, err
	}

	return strings.Split(s, "\r\n"), nil
}