package odata

import (
	"net/http"
	"github.com/amsokol/go-odata/v4"
	"errors"
)

type Request struct {
	Version string
	MaxVersion string
	Data interface{}
}

func ParseRequest(request *http.Request, baseUrl string) (oRequest *Request, err error) {
	//	var regexp = regexp.MustCompile(`^(\w+)(\('(.+)'\))?$`)
	//	return Request{Resource: "qwerty", Others: regexp.FindStringSubmatch(resources)}, err

	oRequest = &Request{
		MaxVersion: request.Header.Get("OData-MaxVersion"),
		Version: request.Header.Get("OData-Version")}

	if len(oRequest.MaxVersion) == 0 {
		// use version 4.0 by default
		oRequest.MaxVersion = v4.ODataVersion
	}

	if len(oRequest.Version) == 0 {
		oRequest.Version = oRequest.MaxVersion
	}

	switch oRequest.Version {
		case v4.ODataVersion:
			oRequest.Data, err = v4.ParseRequestData(request, baseUrl)
		default:
			err = errors.New("Unsupported OData version")
	}

	if err != nil {
		oRequest = nil
	}
	return oRequest, err
}
