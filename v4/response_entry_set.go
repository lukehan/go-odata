package v4

type ResponseEntrySet struct {
	ResponseHeader

	ODataNextLink string `json:"@odata.nextLink,omitempty"` // example = "http://services.odata.org/V4/TripPinService/People?%24skiptoken=8"

	Entries interface{} `json:"value"`
}
type EntryAnnotation struct {
	ODataContext string `json:"@odata.context,omitempty"` // example = "http://services.odata.org/V4/TripPinService/$metadata#People/$entity"
	ODataId string `json:"@odata.id,omitempty"` // example = "http://services.odata.org/V4/TripPinService/People('russellwhyte')"
	ODataETag string `json:"@odata.etag,omitempty"` // example = "W/\"08D256BD91FA2AEF\""
	ODataEditLink string `json:"@odata.editLink,omitempty"` // example = "http://services.odata.org/V4/TripPinService/People('russellwhyte')"
}

func (e *EntryAnnotation) SetId(resourceUrl string, id string) {
	e.ODataId = resourceUrl +  "('" + id + "')"
}

func (e *EntryAnnotation) SetEditLink(resourceUrl string, id string) {
	e.ODataEditLink = resourceUrl +  "('" + id + "')"
}