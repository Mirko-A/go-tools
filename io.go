package io

import (
	"io"
	"os"
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