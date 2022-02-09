package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

/*This var is a pointer towards template.Template that is a
pointer to help process the html.*/
var tpl *template.Template

/*This init function, once it's initialised, makes it so that each html file
in the templates folder is parsed i.e. they all get looked through once and
then stored in the memory ready to go when needed*/
// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*html"))
// }

var (
	ArtistID              int
	ArtistImage           string
	ArtistName            string
	ArtistMembers         []string
	ArtistCreationDate    int
	ArtistFirstAlbum      string
	ArtistLocations       []string
	ArtistConcertDates    []string
	ArtistsDatesLocations map[string][]string
)

// type TotalInfo struct {
// 	ArtistID              int                 `json:"id"`
// 	ArtistImage           string              `json:"image"`
// 	ArtistName            string              `json:"name"`
// 	ArtistMembers         []string            `json:"members"`
// 	ArtistCreationDate    int                 `json:"creationDate"`
// 	ArtistFirstAlbum      string              `json:"firstAlbum"`
// 	ArtistLocations       []string            `json:"locations"`
// 	ArtistConcertDates    []string            `json:"concertDates"`
// 	ArtistsDatesLocations map[string][]string `json:"datesLocations"`
// }

// var Totale []TotalInfo

type InfoAll struct {
	Artists   []OrigArtists
	Dates     []Dates
	Locations []Locations
	Relations []Relations
}

type OrigArtists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var Info InfoAll
var Datos IndexDates
var Connection IndexRelations
var Places IndexLocations
var Artistes OrigArtists

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

// type Relations struct {
// 	Index []struct {
// 		ID             int
// 		DatesLocations map[string][]string
// 	}
// }

type IndexRelations struct {
	Relations []Relations
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func main() {

	if UnmarshalArtistData() != nil || UnmarshalDatesData() != nil || UnmarshalRelationsData() != nil || UnmarshallLocationsData() != nil {

		UnmarshalArtistData()
		UnmarshallLocationsData()
		UnmarshalRelationsData()
		UnmarshalDatesData()

		// for i := range Info.Artists {
		// 	var gd InfoAll
		// 	gd.Artists[i].ID = i + 1
		// 	gd.Artists[i].Image = Info.Artists[i].Image
		// 	gd.Artists[i].Name = Info.Artists[i].Name
		// 	gd.Artists[i].Members = Info.Artists[i].Members
		// 	gd.Artists[i].CreationDate = Info.Artists[i].CreationDate
		// 	gd.Artists[i].FirstAlbum = Info.Artists[i].FirstAlbum
		// 	gd.Locations[i].Locations = Info.Locations[i].Locations
		// 	gd.Dates[i].Dates = Info.Dates[i].Dates
		// 	gd.Relations[i].DatesLocations = Info.Relations[i].DatesLocations
		// 	// = append(Totale, gd)

	}
	for i := 0; i < 52; i++ {
		fmt.Println(Info.Locations[i].Locations)
	}
}

func UnmarshalArtistData() error {

	responseArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get Artists info from API")
	}
	defer responseArtists.Body.Close()

	responseArtistsData, err := ioutil.ReadAll(responseArtists.Body)
	if err != nil {
		panic("Couldn't read data for Artists!")
	}

	// var responseObjectArtists Artists

	json.Unmarshal(responseArtistsData, &Info.Artists)
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

	// var ResponseObjectRelations Relations

	json.Unmarshal(responseData, &Connection)
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

	// var ResponseObjectDates Dates
	json.Unmarshal(responseDatesData, &Datos)
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

	// var ResponseObjectLocations Locations
	json.Unmarshal(responseLocationsData, &Places)
	Info.Locations = Places.Locations

	return nil
}

// func TotalData() {
// 	// UnmarshalArtistData()
// 	// UnmarshallLocationsData()
// 	// UnmarshalRelationsData()
// 	// UnmarshalDatesData()
// 	// for i := range Artists {
// 	// 	var gd TotalInfo
// 	// 	gd.ArtistID = i + 1
// 	// 	gd.ArtistImage = Artists[i].Image
// 	// 	gd.ArtistName = Artists[i].Name
// 	// 	gd.ArtistMembers = Artists[i].Members
// 	// 	gd.ArtistCreationDate = Artists[i].CreationDate
// 	// 	gd.ArtistFirstAlbum = Artists[i].FirstAlbum
// 	// 	gd.ArtistLocations = responseObjectLocations.Locations[i].Locations
// 	// 	gd.ArtistConcertDates = responseObjectDates.Dates[i].Dates
// 	// 	gd.ArtistsDatesLocations = responseObjectRelations.Relations[i].DatesLocations
// 	// 	Totale = append(Totale, gd)

// 	// }
// }

func Requests() {

	http.HandleFunc("/", index)
	http.HandleFunc("/info", artistInfo)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started on: http://localhost:8080")
}

func index(w http.ResponseWriter, r *http.Request) {

	//-------------Create a struct to hold unmarshalled data-----------
	// var IT TI

	if r.URL.Path != "/" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "index.html", nil)
	}
}

func artistInfo(w http.ResponseWriter, r *http.Request) {

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		panic("Couldn't get the relations data!")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Couldn't read data for the Artists")
	}

	var responseObject Relations

	json.Unmarshal(responseData, &responseObject)

	if r.URL.Path != "/info" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "info.html", Connection.Relations)
	}

}
