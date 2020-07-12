package fileUtils

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func FindProjectRoot(rootName string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	elems := strings.Split(dir, "/")
	i := len(elems) - 1

	for i >= 0 && elems[i] != rootName {
		i--
	}

	if i < 0 {
		return "", errors.New(fmt.Sprintf("root not found, cannot find folder %s in current dir %s", rootName, dir))
	}

	return strings.Join(elems[:i + 1], "/"), nil
}
