package fileUtils

import (
	"os"
	"path/filepath"
)

// Open a file and create it if it doesn't exist already
func Upsert(path string) (*os.File, error) {
	file := &os.File{}

	// Check if file exist
	if _, err := os.Stat(path); err == nil {
		// Try opening the file
		file, err = os.Open(path)

		// Couldn't open the file
		if err != nil {
			return nil, err
		}
	} else if os.IsNotExist(err) {
		// Create path if not exist
		if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
			return nil, err
		}

		// File doesn't exist, so we create it
		file, err = os.Create(path)

		// Couldn't create the file
		if err != nil {
			return nil, err
		}
	} else {
		// Unexpected error occurred while checking for file existence
		return nil, err
	}

	// Everything went fine, we can return a safe file object
	return file, nil
}
