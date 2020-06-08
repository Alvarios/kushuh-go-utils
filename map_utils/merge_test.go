package map_utils

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	// Basic merge with 1-level deepness.
	expected := "map[key1:value1 key2:233 key3:value3]"

	r, err := Merge(
		map[string]interface{}{"key1": "value1", "key2": 233},
		map[string]interface{}{"key3": "value3"},
	)

	if err != nil {
		t.Errorf("Error during test unit 1 : %s", err.Error())
	}

	if fmt.Sprintf("%v", r) != expected {
		t.Errorf("Expected %s, got %v", expected, r)
	}

	// Key merging
	expected = "map[key1:value1 key2:alma key3:value3]"

	r, err = Merge(
		map[string]interface{}{"key1": "value1", "key2": 233},
		map[string]interface{}{"key3": "value3", "key2": "alma"},
	)

	if err != nil {
		t.Errorf("Error during test unit 2 : %s", err.Error())
	}

	if fmt.Sprintf("%v", r) != expected {
		t.Errorf("Expected %s, got %v", expected, r)
	}

	// Key replacement
	expected = "map[key1:value1 key2:map[key21:value21 key22:100] key3:value3]"

	r, err = Merge(
		map[string]interface{}{"key1": "value1", "key2": map[string]interface{}{"key21": "value21"}},
		map[string]interface{}{"key3": "value3", "key2": map[string]interface{}{"key22": 100}},
	)

	if err != nil {
		t.Errorf("Error during test unit 2 : %s", err.Error())
	}

	if fmt.Sprintf("%v", r) != expected {
		t.Errorf("Expected %s, got %v", expected, r)
	}

	// Key replacement wth object on 1 side.
	expected = "map[key1:value1 key2:100 key3:value3]"

	r, err = Merge(
		map[string]interface{}{"key1": "value1", "key2": map[string]interface{}{"key21": "value21"}},
		map[string]interface{}{"key3": "value3", "key2": 100},
	)

	if err != nil {
		t.Errorf("Error during test unit 2 : %s", err.Error())
	}

	if fmt.Sprintf("%v", r) != expected {
		t.Errorf("Expected %s, got %v", expected, r)
	}
}
