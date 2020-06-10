package fileUtils

import (
	"strings"
	"testing"
)

func TestFindProjectRoot(t *testing.T) {
	path, err := FindProjectRoot("kushuh-go-utils")
	if err != nil {
		t.Errorf("Error getting project root : %s", err.Error())
	}

	actualRoot := "/Users/kushuh/go/src/kushuh-go-utils"
	if path != actualRoot {
		t.Errorf("Wrong root path returned : expected %s, got %s", actualRoot, path)
	}

	path, err = FindProjectRoot("kushuh-go")
	if err == nil {
		t.Error("No error when providing non existing root name.")
	} else if !strings.HasPrefix(err.Error(), "Root not found") {
		t.Errorf("Unexpected error : %s", err.Error())
	}
}
