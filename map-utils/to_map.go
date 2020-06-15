package mapUtils

import (
	"encoding/json"
	"github.com/Alvarios/kushuh-go-utils/router-utils/responses"
)

func ToMap(v interface{}) (map[string]interface{}, *responses.Error) {
	var output map[string]interface{}

	jsonString, err := json.Marshal(v)
	if err != nil {
		return nil, &responses.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	err = json.Unmarshal(jsonString, &output)
	if err != nil {
		return nil, &responses.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	return output, nil
}
