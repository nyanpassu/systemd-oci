package utils

import (
	"os"
)

// FileExists .
func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func EnsureDirExists(path string) error {
	exists, err := FileExists(path)
	if err != nil {
		return err
	}
	if !exists {
		return os.Mkdir(path, 0644)
	}
	return nil
}
