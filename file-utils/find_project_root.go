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

	elems, i := strings.Split(dir, "/"), 0
	for elems[i] != rootName && i < (len(elems) - 1) {
		i++
	}

	if i == len(elems) - 1 && elems[i] != rootName {
		return "", errors.New(fmt.Sprintf("Root not found : cannot find folder %s in current dir %s", rootName, dir))
	}

	return strings.Join(elems[:i + 1], "/"), nil
}
