package v4
import (
	"net/http"
	"net/url"
)

type ResponseEntrySet struct {
	ResponseHeader

	ODataNextLink string `json:"@odata.nextLink,omitempty"` // example = "http://services.odata.org/V4/TripPinService/People?%24skiptoken=8"

	Entries interface{} `json:"value"`
}
type EntryAnnotation struct {
	ODataId string `json:"@odata.id,omitempty"` // example = "http://services.odata.org/V4/TripPinService/People('russellwhyte')"
	ODataETag string `json:"@odata.etag,omitempty"` // example = "W/\"08D256BD91FA2AEF\""
	ODataEditLink string `json:"@odata.editLink,omitempty"` // example = "http://services.odata.org/V4/TripPinService/People('russellwhyte')"
}

func (e *EntryAnnotation) SetId(id string, request *http.Request) {
	url := url.URL{
		Scheme: request.URL.Scheme,
		Opaque: request.URL.Opaque,
		Host: request.URL.Host,
		Path: request.URL.Path}
	e.ODataId = url.String() +  "('" + id + "')"
}

func (e *EntryAnnotation) SetEditLink(id string, request *http.Request) {
	url := url.URL{
		Scheme: request.URL.Scheme,
		Opaque: request.URL.Opaque,
		Host: request.URL.Host,
		Path: request.URL.Path}
	e.ODataEditLink = url.String() +  "('" + id + "')"
}