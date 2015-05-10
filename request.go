package odata

import (
	"net/http"
	"github.com/amsokol/go-odata/v4"
)

type Request struct {
	Version string
	MaxVersion string
}

func ParseRequest(request *http.Request) (oRequest *Request, err error) {
	oRequest = &Request{
		MaxVersion: request.Header.Get(v4.HEADER_ODataMaxVersion),
		Version: request.Header.Get(v4.HEADER_ODataVersion)}

	if len(oRequest.MaxVersion) == 0 {
		// use version 4.0 by default
		oRequest.MaxVersion = v4.ODataVersion
	}

	if len(oRequest.Version) == 0 {
		oRequest.Version = oRequest.MaxVersion
	}

	return oRequest, nil
}
