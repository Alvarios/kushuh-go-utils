package sliceUtils

import "testing"

func TestCompare(t *testing.T) {
	r := Compare([]int{1,2,3,4,5}, []int{1,2,3,4,5})
	if !r {
		t.Error("two similar arrays of int were labelled as non equal")
		return
	}

	r = Compare(nil, nil)
	if !r {
		t.Error("two nil arrays were labelled as non equal")
		return
	}

	r = Compare([]interface{}{"string", 123}, []interface{}{"string", 123})
	if !r {
		t.Error("two similar arrays of interfaces were labelled as non equal")
		return
	}

	r = Compare([]interface{}{"string", 123}, []int{456, 123})
	if r {
		t.Error("two different arrays of different types were labelled as equal")
		return
	}

	r = Compare([]int{123, 123}, []int{456, 123})
	if r {
		t.Error("two different arrays of same types were labelled as equal")
		return
	}
}
