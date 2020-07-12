package fileUtils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// For upsert and delete functions
func TestIUD(t *testing.T) {
	root, err := FindProjectRoot("file-utils")
	if err != nil {
		t.Errorf("unexpected error : %s", err.Error())
		return
	}

	path := fmt.Sprintf("%s/test.txt", root)
	_ = Delete(path)

	file, err := Upsert(path)
	if err != nil {
		t.Errorf("unable to insert file : %s", err.Error())
		return
	}

	_, err = file.WriteString("hello world")
	if err != nil {
		t.Errorf("unable to write file : %s", err.Error())
		return
	}

	err = file.Close()
	if err != nil {
		t.Errorf("unable to close file : %s", err.Error())
		return
	}

	file, err = Upsert(path)
	if err != nil {
		t.Errorf("unable to update file : %s", err.Error())
		return
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("unable to read file : %s", err.Error())
		return
	}

	if string(b) != "hello world" {
		t.Errorf("unexpected file content : expected \"hello world\", got %s", string(b))
		return
	}

	err = file.Close()
	if err != nil {
		t.Errorf("unable to close file (2) : %s", err.Error())
		return
	}

	err = Delete(path)
	if err != nil {
		t.Errorf("unable to delete file : %s", err.Error())
		return
	}

	path = fmt.Sprintf("%s/log/test.txt", root)
	_ = Delete(path)

	file, err = Upsert(path)
	if err != nil {
		t.Errorf("unable to insert file with incomplete path : %s", err.Error())
		return
	}

	err = file.Close()
	if err != nil {
		t.Errorf("unable to close file (3) : %s", err.Error())
		return
	}

	err = Delete(path)
	if err != nil {
		t.Errorf("unable to delete file : %s", err.Error())
		return
	}

	path = fmt.Sprintf("%s/log", root)
	err = Delete(path)
	if err != nil {
		t.Errorf("unable to delete folder : %s", err.Error())
		return
	}
}
