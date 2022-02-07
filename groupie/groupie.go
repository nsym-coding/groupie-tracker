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
	ArtistID              []int
	ArtistImage           []string
	ArtistName            []string
	ArtistMembers         [][]string
	ArtistCreationDate    []int
	ArtistFirstAlbum      []string
	ArtistLocations       [][]string
	ArtistConcertDates    [][]string
	ArtistsDatesLocations []map[string][]string
	Artistes              []Artists
	//InfoAll               []TotalInfo
	Places Locations
	Days   Dates
	Dlocs  Relations
)

var TotalInfo struct {
	ArtistID              []int
	ArtistImage           []string
	ArtistName            []string
	ArtistMembers         [][]string
	ArtistCreationDate    []int
	ArtistFirstAlbum      []string
	ArtistLocations       [][]string
	ArtistConcertDates    [][]string
	ArtistsDatesLocations []map[string][]string
}

// type TotalInfo []struct {
// 	Totale
// 	ArtistsDatesLocations map[string][]string
// }

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Dates struct {
	Dates []dates `json:"index"`
}

type dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Locations struct {
	Locations []locations `json:"index"`
}

type locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Relations struct {
	Index []struct {
		ID             int
		DatesLocations map[string][]string
	}
}

// type Relations struct {
// 	Relations []relations
// }

// type relations struct {
// 	ID             int                 `json:"id"`
// 	DatesLocations map[string][]string `json:"datesLocations"`
// }

func main() {

	UnmarshalArtistData()

	for i := 0; i < len(TotalInfo.ArtistCreationDate); i++ {
		fmt.Println(TotalInfo.ArtistID[i])
		fmt.Println(TotalInfo.ArtistName[i])
		fmt.Println(TotalInfo.ArtistCreationDate[i])
		fmt.Println(TotalInfo.ArtistFirstAlbum[i])
		fmt.Println(TotalInfo.ArtistMembers[i])
		fmt.Println(TotalInfo.ArtistConcertDates[i])
		fmt.Println(TotalInfo.ArtistLocations[i])
		fmt.Println()

	}

	//FillStruct()

}

func UnmarshalArtistData() {

	responseArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get Artists info from API")
	}
	defer responseArtists.Body.Close()

	responseArtistsData, err := ioutil.ReadAll(responseArtists.Body)
	if err != nil {
		panic("Couldn't read data for Artists!")
	}

	var responseObjectArtists []Artists

	json.Unmarshal(responseArtistsData, &responseObjectArtists)

	for i := 0; i < len(responseObjectArtists); i++ {
		TotalInfo.ArtistFirstAlbum = append(TotalInfo.ArtistFirstAlbum, responseObjectArtists[i].FirstAlbum)
		//fmt.Println(TotalInfo.ArtistFirstAlbum)
	}

	for i := 0; i < len(responseObjectArtists); i++ {
		TotalInfo.ArtistID = append(TotalInfo.ArtistID, responseObjectArtists[i].ID)
		//fmt.Println(TotalInfo.ArtistID)
	}

	for i := 0; i < len(responseObjectArtists); i++ {
		TotalInfo.ArtistImage = append(TotalInfo.ArtistImage, responseObjectArtists[i].Image)

	}

	for i := 0; i < len(responseObjectArtists); i++ {
		TotalInfo.ArtistMembers = append(TotalInfo.ArtistMembers, responseObjectArtists[i].Members)
		//fmt.Println(TotalInfo.ArtistMembers[i])

	}

	for i := 0; i < len(responseObjectArtists); i++ {
		TotalInfo.ArtistCreationDate = append(TotalInfo.ArtistCreationDate, responseObjectArtists[i].CreationDate)
		//fmt.Println(TotalInfo.ArtistCreationDate)

	}

	for i := 0; i < len(responseObjectArtists); i++ {
		TotalInfo.ArtistName = append(TotalInfo.ArtistName, responseObjectArtists[i].Name)
	}

	responseRelations, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		panic("Couldn't get the relations data!")
	}

	responseData, err := ioutil.ReadAll(responseRelations.Body)
	if err != nil {
		panic("Couldn't read data for the Relations")
	}

	var responseObjectRelations Relations

	json.Unmarshal(responseData, &responseObjectRelations)

	// for i := 0; i < 52; i++ {
	// 	for y, v := range responseObjectRelations.Index[i].DatesLocations {

	// 		fmt.Println(y, v)

	// 	}
	// 	fmt.Println()
	// }

	responseDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		panic("Couldn't get Dates info from the API!")
	}
	defer responseDates.Body.Close()

	responseDatesData, err := ioutil.ReadAll(responseDates.Body)
	if err != nil {
		panic("Couldn't read data for Dates")
	}

	var responseObjectDates Dates
	json.Unmarshal(responseDatesData, &responseObjectDates)

	for i := 0; i < len(responseObjectDates.Dates); i++ {
		TotalInfo.ArtistConcertDates = append(TotalInfo.ArtistConcertDates, responseObjectDates.Dates[i].Dates)
	}

	responseLocations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		panic("Couldn't get Location info from API")
	}
	defer responseLocations.Body.Close()

	responseLocationsData, err := ioutil.ReadAll(responseLocations.Body)
	if err != nil {
		panic("Couldn't read data for Locations!")
	}

	var responseObjectLocations Locations
	json.Unmarshal(responseLocationsData, &responseObjectLocations)

	//fmt.Println(responseObjectLocations.Locations[0].Locations)

	for i := 0; i < len(responseObjectLocations.Locations); i++ {
		TotalInfo.ArtistLocations = append(TotalInfo.ArtistLocations, responseObjectLocations.Locations[i].Locations)
	}

}

// func FillStruct() {

// 	for i := range Artistes {

// 		var add TotalInfo

// 		add.ArtistID = Artistes[i].ID
// 		add.ArtistImage = Artistes[i].Image
// 		add.ArtistName = Artistes[i].Name
// 		add.ArtistMembers = Artistes[i].Members
// 		add.ArtistCreationDate = Artistes[i].CreationDate
// 		add.ArtistFirstAlbum = Artistes[i].FirstAlbum
// 		add.ArtistLocations = Places.Locations[i].Locations
// 		add.ArtistConcertDates = Days.Dates[i].Dates
// 		add.ArtistsDatesLocations = Dlocs.Relations[i].DatesLocations

// 		InfoAll = append(InfoAll, add)
// 	}

// }

func Requests() {

	http.HandleFunc("/", index)
	http.HandleFunc("/info", artistInfo)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started on: http://localhost:8080")
}

func index(w http.ResponseWriter, r *http.Request) {

	//-------------Create a struct to hold unmarshalled data-----------

	var TotalInfo struct {
		ArtistID           []int
		ArtistImage        string
		ArtistName         []string
		ArtistMembers      [][]string
		ArtistCreationDate []int
		ArtistFirstAlbum   []string
		ArtistLocations    [][]string
		ArtistConcertDates [][]string
	}

	if r.URL.Path != "/" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "index.html", TotalInfo)
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

		tpl.ExecuteTemplate(w, "info.html", ArtistsDatesLocations)
	}

}
