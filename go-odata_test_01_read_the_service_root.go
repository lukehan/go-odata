package odata

import (
	"net/http"
	"github.com/emicklei/go-restful"
	"github.com/amsokol/go-odata/v4"
)

func ReadTheServiceRoot(request *restful.Request, response *restful.Response) {
	oRequest, err := ParseRequest(request.Request, serviceRoot)

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

	oRequestData, ok := oRequest.Data.(*v4.RequestData)
	if !ok {
		response.WriteErrorString(http.StatusInternalServerError, "RequestData version 4.0 can't be converted to v4.RequestData")
		return
	}

	if oRequestData.Command != v4.CmdReadServiceRoot {
		response.WriteErrorString(http.StatusInternalServerError, "Wrong test method")
		return
	}

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

	var oResponse *v4.Response
	oResponse, err = v4.PrepareResponse(response.Header(), items)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteEntity(oResponse)
}