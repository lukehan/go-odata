package v4

import (
	"github.com/emicklei/go-restful"
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

func WriteResponse(response *restful.Response, oResponse Response) error {
	response.AddHeader("OData-Version", ODataVersion)
	return response.WriteEntity(oResponse)
}