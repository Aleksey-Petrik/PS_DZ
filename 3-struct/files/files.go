package files

import (
	"errors"
	"os"
	"strings"
)

func ReadFile(filePath string) ([]byte, error) {
	if !IsJson(filePath) {
		return []byte{}, errors.New("Файл не является JSON!")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}

	return content, nil
}

func IsJson(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	if info.IsDir() {
		return false
	}

	if !strings.HasSuffix(strings.ToUpper(info.Name()), ".JSON") {
		return false
	}

	return true
}
