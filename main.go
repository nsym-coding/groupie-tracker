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
func init() {
	tpl = template.Must(template.ParseGlob("templates/*html"))
}

var (
	ArtistID              []int
	ArtistImage           []string
	ArtistName            []string
	ArtistMembers         []string
	ArtistCreationDate    []int
	ArtistFirstAlbum      []string
	ArtistLocations       []string
	ArtistConcertDates    []string
	ArtistsDatesLocations map[string][]string
)

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

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func ArtistsData(string) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/Artists")
	if err != nil {
		panic("Couldn't get info for Artists!")
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Couldn't read data for Artists!")
	}

	var responseObject Artists
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject); i++ {

		ArtistID = append(ArtistID, responseObject[i].ID)
	}

	for i := 0; i < len(responseObject); i++ {

		ArtistImage = append(ArtistImage, responseObject[i].Image)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistName = append(ArtistName, responseObject[i].Name)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistMembers = append(ArtistMembers, responseObject[i].Members...)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistCreationDate = append(ArtistCreationDate, responseObject[i].CreationDate)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistFirstAlbum = append(ArtistFirstAlbum, responseObject[i].FirstAlbum)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistLocations = append(ArtistLocations, responseObject[i].Locations)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistConcertDates = append(ArtistConcertDates, responseObject[i].ConcertDates)
	}
	//fmt.Println(ArtistName[0])

}

func main() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get info for Artists!")
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Couldn't read data for Artists!")
	}

	var responseObject Artists
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject); i++ {

		ArtistID = append(ArtistID, responseObject[i].ID)
	}

	for i := 0; i < len(responseObject); i++ {

		ArtistImage = append(ArtistImage, responseObject[i].Image)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistName = append(ArtistName, responseObject[i].Name)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistMembers = append(ArtistMembers, responseObject[i].Members...)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistCreationDate = append(ArtistCreationDate, responseObject[i].CreationDate)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistFirstAlbum = append(ArtistFirstAlbum, responseObject[i].FirstAlbum)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistLocations = append(ArtistLocations, responseObject[i].Locations)
	}
	for i := 0; i < len(responseObject); i++ {

		ArtistConcertDates = append(ArtistConcertDates, responseObject[i].ConcertDates)
	}

	fmt.Println(ArtistImage[0:5])
	requests()

}

func requests() {
	fs := http.FileServer(http.Dir("./templates"))

	http.Handle("/", fs)
	http.HandleFunc("/index.html", index)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started on: http://localhost:8080")
}
func index(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get info for Artists!")
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Couldn't read data for Artists!")
	}

	var responseObject Artists
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject); i++ {

		ArtistImage = append(ArtistImage, responseObject[i].Image)
	}
	if r.URL.Path != "/" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "index.html", ArtistImage)
	}
}
