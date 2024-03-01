package util

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/crypto"
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

func RandomPrivateKey() ([]byte, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return crypto.FromECDSA(privateKey), nil
}
