package cbErrors

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Pretty(err error) (string, error) {
	errS := err.Error()
	errType := strings.Split(errS, " | ")[0]
	errJSON := errS[len(errType) + 3:]

	var errJSONU map[string]interface{}
	e := json.Unmarshal([]byte(errJSON), &errJSONU)
	if e != nil {
		return "", fmt.Errorf("%s\n\n%s\n\n", e.Error(), errJSON)
	}

	errJSONI, e := json.MarshalIndent(errJSONU, "", "    ")
	if e != nil {
		return "", e
	}

	return fmt.Sprintf("%s\n%s", errType, errJSONI), nil
}