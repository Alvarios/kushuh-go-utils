package mapUtils

import (
	"fmt"
	"testing"
)

func TestAssignPath(t *testing.T) {
	r := AssignPath(map[string]interface{}{}, "key1.key11", 12)
	expected := "map[key1:map[key11:12]]"
	if fmt.Sprint(r) != expected {
		t.Errorf("unexpected value : expected %s, got %v", expected, r)
	}

	r = AssignPath(r, "key1.key21", "a string")
	expected = "map[key1:map[key11:12 key21:a string]]"
	expected2 := "map[key1:map[key21:a string key11:12]]"
	if fmt.Sprint(r) != expected && fmt.Sprint(r) != expected2  {
		t.Errorf("unexpected value : expected %s, got %v", expected, r)
	}

	r = AssignPath(r, "key1.key11", "a string")
	expected = "map[key1:map[key11:a string key21:a string]]"
	expected2 = "map[key1:map[key21:a string key11:a string]]"
	if fmt.Sprint(r) != expected && fmt.Sprint(r) != expected2  {
		t.Errorf("unexpected value : expected %s, got %v", expected, r)
	}

	r = AssignPath(r, "key1", "a string")
	expected = "map[key1:a string]"
	if fmt.Sprint(r) != expected  {
		t.Errorf("unexpected value : expected %s, got %v", expected, r)
	}
}
