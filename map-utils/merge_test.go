package mapUtils

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

	// Documentation test.
	expected = "map[key1:value1 key2:value2 key3:map[key31:map[ultra-nested:[]] key32:11] key4:value4]"

	original := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"key21": 10,
			"key22": true,
		},
		"key3": map[string]interface{}{
			"key31": map[string]interface{}{
				"ultra-nested": []string{},
			},
			"key32": false,
		},
	}

	replacement1 := map[string]interface{}{
		"key2": "value2",
		"key3": map[string]interface{}{
			"key32": 11,
		},
	}

	replacement2 := map[string]interface{}{
		"key4": "value4",
	}

	r, err = Merge(original, replacement1, replacement2)

	if err != nil {
		t.Errorf("Error during test unit 2 : %s", err.Error())
	}

	if fmt.Sprintf("%v", r) != expected {
		t.Errorf("Expected %s, got %v", expected, r)
	}
}
