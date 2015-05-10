package v4

type ResponseHeader struct {
	ODataContext string `json:"@odata.context,omitempty"` // example = "http://services.odata.org/V4/TripPinService/$metadata"
}
