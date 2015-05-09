package v4

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
