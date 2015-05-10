package v4
import (
	"encoding/json"
)

type ResponseError struct {
	Error Error `json:"error"`
}

type Error struct {
	Code string `json:"code"`
	Message string `json:"message"`
	InnerError *InnerError `json:"innererror,omitempty"`
}

type InnerError struct {
	Message string `json:"message"`
	Type string `json:"type"`
	StackTrace string `json:"stacktrace"`
}

func (e *ResponseError) String() string {
	j, _ := json.Marshal(e)
	return string(j)
}