package groupie

import (
	"encoding/json"
	"io/ioutil"
	//"log"
	"net/http"
	"fmt"
)


var Information []TotalInfo

type TotalInfo struct {
	Artists   []Artists
	Dates     []Dates
	Locations []Locations
	Relations []Relations
}

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type IndexDates struct {
	Dates []Dates `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type IndexLocations struct {
	Locations []Locations `json:"index"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type IndexRelations struct {
	Relations []Relations `json:"index"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var (
	 Info TotalInfo
	 Datos IndexDates
	 Connection IndexRelations
	 Places IndexLocations
)
func main() {
	UnmarshalArtistData()
	UnmarshalDatesData()
	UnmarshallLocationsData()
}

func UnmarshalArtistData() error {
	//--------------Unmarshall Artists-------------------

	responseArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get Artists info from API")
	}
	defer responseArtists.Body.Close()

	responseArtistsData, err := ioutil.ReadAll(responseArtists.Body)
	if err != nil {
		panic("Couldn't read data for Artists!")
	}
	//--------------Unmarshalled all Artist data directly to Info variable-------------------
	json.Unmarshal(responseArtistsData, &Info.Artists)
	fmt.Println(Info.Artists[10])
	return nil
}

func UnmarshalRelationsData() error {
	responseRelations, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		panic("Couldn't get the relations data!")
	}

	responseData, err := ioutil.ReadAll(responseRelations.Body)
	if err != nil {
		panic("Couldn't read data for the Relations")
	}

	//--------------Unmarshalled relation data directly to Connection variable-------------------
	json.Unmarshal(responseData, &Connection)
	Info.Relations = Connection.Relations
	// fmt.Println(Info.Relations)
	return nil
}

func UnmarshalDatesData() error {
	responseDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		panic("Couldn't get Dates info from the API!")
	}
	defer responseDates.Body.Close()

	responseDatesData, err := ioutil.ReadAll(responseDates.Body)
	if err != nil {
		panic("Couldn't read data for Dates")
	}

	//--------------Unmarshalled dates data directly to Datos variable-------------------
	json.Unmarshal(responseDatesData, &Datos)
	Info.Dates = Datos.Dates
	return nil
}

func UnmarshallLocationsData() error {
	responseLocations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		panic("Couldn't get Location info from API")
	}
	defer responseLocations.Body.Close()

	responseLocationsData, err := ioutil.ReadAll(responseLocations.Body)
	if err != nil {
		panic("Couldn't read data for Locations!")
	}
	//--------------Unmarshalled locations data directly to Places variable-------------------
	json.Unmarshal(responseLocationsData, &Places)
	Info.Locations = Places.Locations

	return nil
}