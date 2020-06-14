package responses

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code %v : %s", e.Code, e.Message)
}