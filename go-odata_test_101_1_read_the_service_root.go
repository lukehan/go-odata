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

	var oResponse *v4.Response
	oResponse, err = ReadTheServiceRootV4(response.Header())
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteEntity(oResponse)
}

func ReadTheServiceRootV4(header http.Header) (*v4.Response, error) {
	items := make([]v4.ResourceDefinition, 0, 6)
	items = append(items,
		v4.ResourceDefinition{Name: "Photos", Kind: v4.KindEntitySet, Url: "Photos"})
	items = append(items,
		v4.ResourceDefinition{Name: "People", Kind: v4.KindEntitySet, Url: "People"})
	items = append(items,
		v4.ResourceDefinition{Name: "Airlines", Kind: v4.KindEntitySet, Url: "Airlines"})
	items = append(items,
		v4.ResourceDefinition{Name: "Airports", Kind: v4.KindEntitySet, Url: "Airports"})
	items = append(items,
		v4.ResourceDefinition{Name: "Me", Kind: v4.KindSingleton, Url: "Me"})
	items = append(items,
		v4.ResourceDefinition{Name: "GetNearestAirport", Kind: v4.KindFunctionImport, Url: "GetNearestAirport"})

	return v4.PrepareResponse(header, items)
}
