package groupie

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

var Information []InfoAll

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

//var Artistes []OrigArtists

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

// func main() {
// 	fmt.Println(Info.Relations[1].DatesLocations[Info.Locations[1].Locations[0]])

// }

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

	json.Unmarshal(responseArtistsData, &Info.Artists)
}

func UnmarshalRelationsData() {

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
	Info.Relations = Connection.Relations
	//fmt.Println(Info.Relations)
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

	json.Unmarshal(responseLocationsData, &Places)
	Info.Locations = Places.Locations

	return nil
}

func Requests() {

	http.Handle("/", http.FileServer(http.Dir("./templates")))

	http.HandleFunc("/index", index)
	http.HandleFunc("/relations", relationsInfo)
	http.HandleFunc("/bandmembers", bandMembers)
	http.HandleFunc("/bandlocations", bandLocations)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started on: http://localhost:8080")
}

func index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/index" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "index.html", Info.Artists)

	}
}
func relationsInfo(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/relations" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		submit := r.FormValue("ChosenBandDL")
		Numsubmit, _ := strconv.Atoi(submit)
		fmt.Println(Numsubmit - 1)

		p := Info.Relations[Numsubmit-1].DatesLocations

		fmt.Println(Numsubmit - 1)

		tpl.ExecuteTemplate(w, "relations.html", p)
	}

}

func bandMembers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bandmembers" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {
		submit := r.FormValue("ChosenBandMembers")
		Numsubmit, _ := strconv.Atoi(submit)
		fmt.Println(Numsubmit - 1)

		p := Info.Artists[Numsubmit-1]

		fmt.Println(Numsubmit - 1)

		tpl.ExecuteTemplate(w, "bandmembers.html", p)
	}
}

func bandLocations(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bandlocations" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)

	} else {
		submit := r.FormValue("ChosenBandLocations")
		Numsubmit, _ := strconv.Atoi(submit)
		fmt.Println(Numsubmit - 1)

		// var styleStart string = "<style> h1 { color: red('"
		// var styleEnd string = "'); } </style>"

		fmt.Fprintln(w, "<h1> LOCATIONS </h1>")
		fmt.Fprintln(w, "<pre>"+strings.Join(Info.Locations[Numsubmit-1].Locations, "\n"))

		fmt.Fprintln(w, "<h1> CREATION DATE </h1>")
		fmt.Fprintln(w, Info.Artists[Numsubmit-1].CreationDate)

		fmt.Fprintln(w, "<h1> FIRST ALBUM </h1>")
		fmt.Fprintln(w, Info.Artists[Numsubmit-1].FirstAlbum)

		fmt.Fprintln(w, " <h1> DATES </h1> ")
		fmt.Fprintln(w, "<pre>"+strings.Join(Info.Dates[Numsubmit-1].Dates, "\n"))

		// fmt.Fprintln(w, "<h1> DATES & LOCATIONS </h1>")
		// for _, value := range Info.Relations {
		// 	for place, date := range value.DatesLocations {
		// 		fmt.Fprintln(w, place[Numsubmit-1], date[Numsubmit-1])
		// 	}
		// }
		fmt.Println(Numsubmit - 1)

		tpl.ExecuteTemplate(w, "bandlocations.html", nil)

	}
}
