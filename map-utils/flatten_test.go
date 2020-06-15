package mapUtils

import (
	"fmt"
	"testing"
)

func TestFlattenKeys(t *testing.T) {
	data := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"key21": "value21",
			"key22": map[string]interface{}{
				"key221": "value221",
			},
		},
	}

	flattened := Flatten(data, "")

	if !(
		fmt.Sprint(flattened) == "map[key1:value1 key2.key21:value21 key2.key22.key221:value221]" ||
		fmt.Sprint(flattened) == "map[key1:value1 key2.key22.key221:value221 key2.key21:value21]" ||
		fmt.Sprint(flattened) == "map[key2.key22.key221:value221 key1:value1 key2.key21:value21]" ||
		fmt.Sprint(flattened) == "map[key2.key22.key221:value221 key2.key21:value21 key1:value1]" ||
		fmt.Sprint(flattened) == "map[key2.key21:value21 key2.key22.key221:value221 key1:value1]" ||
		fmt.Sprint(flattened) == "map[key2.key21:value21 key1:value1 key2.key22.key221:value221]") {
		t.Errorf(
			"wrong flattened map : expected something like map[[key1 value1] [key2.key21 value21] [key2.key22.key221 value221]], got %s",
			flattened,
		)
	}
}
