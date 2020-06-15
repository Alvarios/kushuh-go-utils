package mapUtils

import (
	"fmt"
)

// Each nested key will be extracted as a top level key.
// eg: {k1 : val1, k2: {k21: val21}} is flattened as {k1: val1, k2.k21: val21}
func Flatten(params map[string]interface{}, parentKey string) map[string]interface{} {
	output := map[string]interface{}{}

	if params == nil {
		return output
	}

	for key, value := range params {
		// Point to a specific key in the document.
		fKey := fmt.Sprintf("%s%s", parentKey, key)

		if mValue, ok := value.(map[string]interface{}); ok {
			// Map value has to be flattened.
			so := Flatten(mValue, fmt.Sprintf("%s.", fKey))

			for sk, sv := range so {
				output[sk] = sv
			}
		} else {
			// Append value without any check.
			output[fKey] = value
		}
	}

	return output
}
