package v4

import "net/http"

type ResponseHeader struct {
	ODataContext string `json:"@odata.context,omitempty"` // example = "http://services.odata.org/V4/TripPinService/$metadata"
}

func (r *ResponseHeader) AddVersion(header http.Header) {
	header.Add("OData-Version", ODataVersion)
}

/*
type Response interface {
	GetResponse(header http.Header) interface{}
}

func AddVersionToHttpHeader(header http.Header) {
	header.Add("OData-Version", ODataVersion)
}
*/