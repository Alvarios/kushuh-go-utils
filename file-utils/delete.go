package fileUtils

import "os"

func Delete(path string) error {
	// Check if file exist
	if _, err := os.Stat(path); err == nil {
		err = os.Remove(path)

		if err != nil {
			return err
		}

		return nil
	} else if os.IsNotExist(err) {
		// File doesn't exist so we don't need to do anything
		return nil
	} else {
		// Unexpected error occurred while checking for file existence
		return err
	}
}
