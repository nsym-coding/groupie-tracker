package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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
	artistID              []int
	artistImage           []string
	artistName            []string
	artistMembers         []string
	artistCreationDate    []int
	artistFirstAlbum      []string
	artistLocations       []string
	artistConcertDates    []string
	artistsDatesLocations map[string][]string
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

func artistsData(string) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get info for artists!")
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Couldn't read data for artists!")
	}

	var responseObject Artists
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject); i++ {

		artistID = append(artistID, responseObject[i].ID)
	}

	for i := 0; i < len(responseObject); i++ {

		artistImage = append(artistImage, responseObject[i].Image)
	}
	for i := 0; i < len(responseObject); i++ {

		artistName = append(artistName, responseObject[i].Name)
	}
	for i := 0; i < len(responseObject); i++ {

		artistMembers = append(artistMembers, responseObject[i].Members...)
	}
	for i := 0; i < len(responseObject); i++ {

		artistCreationDate = append(artistCreationDate, responseObject[i].CreationDate)
	}
	for i := 0; i < len(responseObject); i++ {

		artistFirstAlbum = append(artistFirstAlbum, responseObject[i].FirstAlbum)
	}
	for i := 0; i < len(responseObject); i++ {

		artistLocations = append(artistLocations, responseObject[i].Locations)
	}
	for i := 0; i < len(responseObject); i++ {

		artistConcertDates = append(artistConcertDates, responseObject[i].ConcertDates)
	}
	//fmt.Println(artistName[0])

}

func main() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get info for artists!")
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Couldn't read data for artists!")
	}

	var responseObject Artists
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject); i++ {

		artistID = append(artistID, responseObject[i].ID)
	}

	for i := 0; i < len(responseObject); i++ {

		artistImage = append(artistImage, responseObject[i].Image)
	}
	for i := 0; i < len(responseObject); i++ {

		artistName = append(artistName, responseObject[i].Name)
	}
	for i := 0; i < len(responseObject); i++ {

		artistMembers = append(artistMembers, responseObject[i].Members...)
	}
	for i := 0; i < len(responseObject); i++ {

		artistCreationDate = append(artistCreationDate, responseObject[i].CreationDate)
	}
	for i := 0; i < len(responseObject); i++ {

		artistFirstAlbum = append(artistFirstAlbum, responseObject[i].FirstAlbum)
	}
	for i := 0; i < len(responseObject); i++ {

		artistLocations = append(artistLocations, responseObject[i].Locations)
	}
	for i := 0; i < len(responseObject); i++ {

		artistConcertDates = append(artistConcertDates, responseObject[i].ConcertDates)
	}

	//fmt.Println(artistImage[0:5])
	requests()

}

func requests() {
	fs := http.FileServer(http.Dir("./templates"))

	http.Handle("/", fs)
	http.HandleFunc("/index.html", index)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(artistImage[0:5])
	if r.URL.Path != "/" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "index.html", artistImage)
	}
}
