package odata

import (
	"github.com/emicklei/go-restful"
	"github.com/amsokol/go-odata/v4"
	"net/http"
)

type Man struct {
	v4.EntryAnnotation

	UserName string `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Emails []string `json:"Emails"`
	AddressInfo AddressInfo `json:"AddressInfo"`
	Gender string `json:"Gender"`
}
type AddressInfo struct {
	Address string `json:"Address"`
	City City `json:"City"`
}

type City struct {
	CountryRegion string `json:"CountryRegion"`
	Name string `json:"Name"`
	Region string `json:"Region"`
}

func ReadAnEntitySet(request *restful.Request, response *restful.Response) {
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

	oResponse := ReadAnEntitySetV4(request)

	response.AddHeader("OData-Version", v4.ODataVersion)
	response.WriteEntity(oResponse)
}

func ReadAnEntitySetV4(request *restful.Request) (r *v4.ResponseEntrySet) {
	entries := []*Man{
		&Man{
			UserName: "russellwhyte",
			FirstName: "Russell",
			LastName: "Whyte",
			Emails: []string{"Russell@example.com", "Russell@contoso.com"},
			AddressInfo : AddressInfo{
				Address: "187 Suffolk Ln.",
				City: City{
					CountryRegion: "United States",
					Name: "Boise",
					Region: "ID"}},
			Gender: "Male"},
		&Man{
			UserName: "javieralfred",
			FirstName: "Javier",
			LastName: "Alfred",
			Emails: []string{"Javier@example.com", "Javier@contoso.com"},
			AddressInfo : AddressInfo{
				Address: "89 Jefferson Way Suite 2",
				City: City{
					CountryRegion: "United States",
					Name: "Portland",
					Region: "WA"}},
			Gender: "Male"}}

	for _,e := range entries  {
		e.SetId(e.UserName, request.Request)
		e.SetEditLink(e.UserName, request.Request)
	}

	return &v4.ResponseEntrySet{Entries:entries}
}
