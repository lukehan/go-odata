package v4
import "net/http"

const (
// Read the service root
	KindEntitySet ResourceKind = "EntitySet"
	KindSingleton ResourceKind = "Singleton"
	KindFunctionImport ResourceKind = "FunctionImport"
)

type ResourceKind string

type Resource struct {
	Name string `json:"name"`
	Kind ResourceKind `json:"kind"`
	Url string `json:"url"`
}

type ResponseResourceList struct {
	ResponseHeader

	Resources []Resource `json:"value"`
}

func (r ResponseResourceList) GetResponse(header http.Header) interface{} {
	header.Add("OData-Version", ODataVersion)
	return nil
}
