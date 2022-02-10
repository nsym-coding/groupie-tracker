package groupie

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

	http.HandleFunc("/", index)
	http.HandleFunc("/info", artistInfo)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started on: http://localhost:8080")
}

func index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "index.html", Info.Artists)

	}
}
func artistInfo(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/info" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		submit := r.FormValue("submit")
		Numsubmit, _ := strconv.Atoi(submit)

		var p []string

		for i := 0; i < 52; i++ {
			if Numsubmit == Info.Artists[i].ID {
				p = Info.Relations[Numsubmit].DatesLocations[Info.Locations[Numsubmit].Locations[0]]
			}
		}

		tpl.ExecuteTemplate(w, "info.html", p)
	}

}
