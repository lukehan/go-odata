package v4

import (
	"net/http"
)

const (
// Read the service root
	KindEntitySet ResourceKind = "EntitySet"
	KindSingleton ResourceKind = "Singleton"
	KindFunctionImport ResourceKind = "FunctionImport"
)

type ResourceKind string

type ResourceDefinition struct {
	Name string `json:"name"`
	Kind ResourceKind `json:"kind"`
	Url string `json:"url"`
}

type Response struct {
	ODataContext string `json:"@odata.context,omitempty"`
	Value interface{} `json:"value"`
}

func PrepareResponse(header http.Header, responseData interface{}) (response *Response, err error) {
	header.Add("OData-Version", ODataVersion)
	return &Response{Value:responseData}, nil
}