package fileUtils

import (
	"fmt"
	"os"
)

func OpenFromProjectRoot(root, path string) (*os.File, error) {
	rootPath, err := FindProjectRoot(root)
	if err != nil {
		return nil, err
	}

	return os.Open(fmt.Sprintf("%s/%s", rootPath, path))
}
