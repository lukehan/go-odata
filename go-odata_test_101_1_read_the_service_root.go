package odata

import (
	"net/http"
	"github.com/emicklei/go-restful"
	"github.com/amsokol/go-odata/v4"
)

func ReadTheServiceRoot(request *restful.Request, response *restful.Response) {
	oRequest, err := ParseRequest(request.Request)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	if oRequest.MaxVersion < v4.ODataVersion {
		response.WriteErrorString(http.StatusBadRequest, "Unsupported OData MaxVersion")
		return
	}

	if oRequest.Version != v4.ODataVersion {
		response.WriteErrorString(http.StatusBadRequest, "Unsupported OData Version")
		return
	}

	oResponse := ReadTheServiceRootV4()

	response.AddHeader(v4.HEADER_ODataVersion, v4.ODataVersion)
	response.WriteEntity(oResponse)
}

func ReadTheServiceRootV4() (r *v4.ResponseResourceList) {
	r = &v4.ResponseResourceList{Resources:[]v4.Resource{
		v4.Resource{Name: "Photos", Kind: v4.KindEntitySet, Url: "Photos"},
		v4.Resource{Name: "People", Kind: v4.KindEntitySet, Url: "People"},
		v4.Resource{Name: "Airlines", Kind: v4.KindEntitySet, Url: "Airlines"},
		v4.Resource{Name: "Airports", Kind: v4.KindEntitySet, Url: "Airports"},
		v4.Resource{Name: "Me", Kind: v4.KindSingleton, Url: "Me"},
		v4.Resource{Name: "GetNearestAirport", Kind: v4.KindFunctionImport, Url: "GetNearestAirport"}}}

	return r
}
