package slack

import (
	httpUtils "github.com/Alvarios/kushuh-go-utils/router-utils/http-utils"
	"net/http"
)

func Send(w string, m string, b map[string]interface{}, a []map[string]interface{}) (*http.Response, error) {
	fm := map[string]interface{}{"text" : m}

	if b != nil {
		fm["blocks"] = b
	}

	if a != nil {
		fm["attachments"] = a
	}

	return httpUtils.Post(w, fm)
}
