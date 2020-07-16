package slack

import (
	httpUtils "github.com/Alvarios/kushuh-go-utils/router-utils/http-utils"
	"net/http"
)

func Send(w string, m string) (*http.Response, error) {
	return httpUtils.Post(w, map[string]interface{}{"text" : m})
}
