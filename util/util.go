package util

import (
	"log/slog"
	"os"
	"path/filepath"
)

func GetRootPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				// at the root
				break
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			slog.Error(err.Error())
			return "", err
		}
		break
	}

	return currentDir, nil
}
