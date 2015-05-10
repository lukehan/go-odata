package odata

import (
	"github.com/emicklei/go-restful"
	"regexp"
	"github.com/amsokol/go-odata/v4"
	"net/http"
)

func GetASingleEntityFromACollection(request *restful.Request, response *restful.Response) {
	resource := request.PathParameter("resource")

	re := regexp.MustCompile("^People\\('(.+)'\\)$")
	sm := re.FindStringSubmatch(resource)

	oResponse := GetASingleEntityFromACollectionV4(sm[1], request)

	if oResponse == nil {
		err := v4.ResponseError{
			Error:v4.Error{
				Code:"NotFound",
				Message:"Entry not found"}}
		response.AddHeader(v4.HEADER_ODataVersion, v4.ODataVersion)
		response.WriteErrorString(http.StatusNotFound, err.String())
		return
	}

	response.AddHeader(v4.HEADER_ODataVersion, v4.ODataVersion)
	response.WriteEntity(oResponse)
}

func GetASingleEntityFromACollectionV4(id string, request *restful.Request) (r *Man) {
	var man *Man;
	switch id {
	case "russellwhyte":
		man = &Man{
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
			Gender: "Male"}
		man.SetId(peopleResourceBaseUrl, man.UserName)
		man.SetEditLink(peopleResourceBaseUrl, man.UserName)
	case "javieralfred":
		man = &Man{
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
			Gender: "Male"}
		man.SetId(peopleResourceBaseUrl, man.UserName)
		man.SetEditLink(peopleResourceBaseUrl, man.UserName)
	}
	return man
}
