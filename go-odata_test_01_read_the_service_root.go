package odata
import (
	"net/http"
	"github.com/emicklei/go-restful"
)

func ReadTheServiceRoot(request *restful.Request, response *restful.Response) {
	oRequest, _ := ParseRequest(request, serviceRoot)

	if oRequest.Command != CmdReadServiceRoot {
		response.WriteErrorString(http.StatusInternalServerError, "Wrong test method")
		return
	}

	items := make([]ResourceDefinition, 0, 6)
	items = append(items,
		ResourceDefinition{Name: "Photos", Kind: KindEntitySet, Url: "Photos"})
	items = append(items,
		ResourceDefinition{Name: "People", Kind: KindEntitySet, Url: "People"})
	items = append(items,
		ResourceDefinition{Name: "Airlines", Kind: KindEntitySet, Url: "Airlines"})
	items = append(items,
		ResourceDefinition{Name: "Airports", Kind: KindEntitySet, Url: "Airports"})
	items = append(items,
		ResourceDefinition{Name: "Me", Kind: KindSingleton, Url: "Me"})
	items = append(items,
		ResourceDefinition{Name: "GetNearestAirport", Kind: KindFunctionImport, Url: "GetNearestAirport"})

	WriteResponse(response, Response{Value: items})
}