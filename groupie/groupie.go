package main

import (
	"encoding/json"
	"fmt"
	//"html/template"
	"io/ioutil"
	//"log"
	"net/http"
)

// /*This var is a pointer towards template.Template that is a
// pointer to help process the html.*/
// var tpl *template.Template

// /*This init function, once it's initialised, makes it so that each html file
// in the templates folder is parsed i.e. they all get looked through once and
// then stored in the memory ready to go when needed*/
// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*html"))
// }

// var (
// 	ArtistID              []int
// 	ArtistImage           []string
// 	ArtistName            []string
// 	ArtistMembers         [][]string
// 	ArtistCreationDate    []int
// 	ArtistFirstAlbum      []string
// 	ArtistLocations       [][]string
// 	ArtistConcertDates    [][]string
// 	ArtistsDatesLocations map[string][]string
// )

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

// type TotalInfo struct {
// 	ArtistID              []int
// 	ArtistImage           []string
// 	ArtistName            []string
// 	ArtistMembers         [][]string
// 	ArtistCreationDate    []int
// 	ArtistFirstAlbum      []string
// 	ArtistLocations       [][]string
// 	ArtistConcertDates    [][]string
// 	ArtistsDatesLocations map[string][]string
// }

type TI struct {
	ArtistID              []int
	ArtistImage           []string
	ArtistName            []string
	ArtistMembers         [][]string
	ArtistCreationDate    []int
	ArtistFirstAlbum      []string
	ArtistLocations       [][]string
	ArtistConcertDates    [][]string
	ArtistsDatesLocations map[string][]string
}

type IT struct {
	IT []TI
}

type Artists []struct {
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


func main(){
	//UnmarshalArtistData()
	UnmarshalDatesLocations()
}

// func UnmarshalArtistData(){
// //--------------Unmarshall Artists-------------------

// 	responseArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
// 	if err != nil {
// 		panic("Couldn't get Artists info from API")
// 	}
// 	defer responseArtists.Body.Close()

// 	responseArtistsData, err := ioutil.ReadAll(responseArtists.Body)
// 	if err != nil {
// 		panic("Couldn't read data for Artists!")
// 	}

// 	var responseObjectArtists Artists
// 	json.Unmarshal(responseArtistsData, &responseObjectArtists)
// //--------------Append Artist ID to variable-------------------
// 	for i := 0; i < len(responseObjectArtists); i++ {
// 		ArtistID = append(ArtistID, responseObjectArtists[i].ID)
// 	}
// //--------------Append Artist Name to variable-------------------

// 	for i := 0; i < len(responseObjectArtists); i++ {
// 		ArtistName = append(ArtistName, responseObjectArtists[i].Name)
// 	}
// //--------------Append Artist Image to variable-------------------

// 	for i := 0; i < len(responseObjectArtists); i++ {
// 		ArtistImage = append(ArtistImage, responseObjectArtists[i].Image)
// 	}

// //--------------Append Artist Members to variable-------------------

// 	for i := 0; i < len(responseObjectArtists); i++ {
// 		ArtistMembers = append(ArtistMembers, responseObjectArtists[i].Members)
// 	}
// //--------------Append Artist Creation Date to variable-------------------

// 	for i := 0; i < len(responseObjectArtists); i++ {
// 		ArtistCreationDate = append(ArtistCreationDate, responseObjectArtists[i].CreationDate)
// 	}

// 	// fmt.Println(ArtistID[0])
// 	// fmt.Println(ArtistName[0])
// 	// fmt.Println(ArtistImage[0])
// 	// fmt.Println(ArtistCreationDate[0])

// }

func UnmarshalDatesLocations() {

//--------------Unmarshall Relations-------------------

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

	
	for _, slice := range responseObjectRelations.Index {
		TotalInfo.ArtistsDatesLocations = append(TotalInfo.ArtistsDatesLocations, slice.DatesLocations)
	}
	
	//fmt.Println(TotalInfo.ArtistsDatesLocations[0])


	
//--------------Unmarshall Dates-------------------

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

	//fmt.Println(TotalInfo.ArtistConcertDates[0])


//--------------Unmarshall Locations-------------------

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

	// for i := 0; i < len(responseObjectLocations.Locations); i++ {
	// 	ArtistLocations = append(ArtistLocations, responseObjectLocations.Locations[i].Locations)
	// }
	for i := 0; i < len(responseObjectLocations.Locations); i++ {
		TotalInfo.ArtistLocations = append(TotalInfo.ArtistLocations, responseObjectLocations.Locations[i].Locations)
	}

	//fmt.Println(TotalInfo.ArtistLocations[0])

}