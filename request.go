package odata

import (
	"strings"
	"net/http"
	"github.com/amsokol/go-odata/v4"
)

// OData commands
const (
	CmdUnknown Command = 0
	CmdReadServiceRoot Command = 1
)

type Command int

type Request struct {
	MaxVersion string
	Command Command
}

func ParseRequest(request *http.Request, baseUrl string) (oRequest Request, err error) {
	//	var regexp = regexp.MustCompile(`^(\w+)(\('(.+)'\))?$`)
	//	return Request{Resource: "qwerty", Others: regexp.FindStringSubmatch(resources)}, err

	oRequest = Request{MaxVersion: request.Header.Get("OData-MaxVersion")}

	if len(oRequest.MaxVersion) == 0 {
		oRequest.MaxVersion = v4.ODataVersion
	}

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