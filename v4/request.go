package v4

import (
	"net/http"
	"strings"
)

const (
	ODataVersion string = "4.0"
)

type Command int

// OData commands
const (
	CmdUnknown Command = 0
	CmdReadServiceRoot Command = 1
)

type RequestData struct {
	Command Command
}

func ParseRequestData(request *http.Request, baseUrl string) (oRequest *RequestData, err error) {
	oRequest = &RequestData{
		Command:CmdUnknown}

	path := strings.Trim(
		strings.TrimPrefix(
			strings.TrimLeft(request.URL.Path, "/"),
			strings.Trim(baseUrl, "/")),
		"/")

	if path = strings.TrimSpace(path); len(path) == 0 {
		oRequest.Command = CmdReadServiceRoot
		return oRequest, nil
	}

	return oRequest, nil
}