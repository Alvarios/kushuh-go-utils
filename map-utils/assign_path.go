package mapUtils

import (
	"strings"
)

func AssignPath(v map[string]interface{}, path string, e interface{}) map[string]interface{} {
	elements := strings.Split(path, ".")

	if len(elements) > 1 {
		if mve, ok := v[elements[0]].(map[string]interface{}); !ok {
			v[elements[0]] = AssignPath(map[string]interface{}{}, strings.Join(elements[1:], "."), e)
		} else {
			v[elements[0]] = AssignPath(mve, strings.Join(elements[1:], "."), e)
		}
	} else {
		v[elements[0]] = e
	}

	return v
}
