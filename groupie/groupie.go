package groupie

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
// port := os.Getenv("PORT")
//     if port == "" {
//         port = "8080" // Default port if not specified (with respect to heroku hosting)
//     }
//     http.ListenAndServe(":"+port, nil)

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

	http.Handle("/css/mystyle.css", http.FileServer(http.Dir("./templates")))

	http.HandleFunc("/", index)
	http.HandleFunc("/bandinfo", bandInfo)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified (with respect to heroku hosting)
	}
	http.ListenAndServe(":"+port, nil)
	//http.ListenAndServe(":8080", nil)
	log.Println("Server started on: http://localhost:8080")
}

func index(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			tpl.ExecuteTemplate(w, "error.html", nil)
			fmt.Fprintln(w, " An internal server error has occurred: ", http.StatusInternalServerError)
			return
		}
	}()

	if r.URL.Path != "/" {

		tpl.ExecuteTemplate(w, "error.html", nil)
		fmt.Fprintln(w, "\n\n\n\n Address not found: wrong address entered, status: ", http.StatusNotFound)

	} else {

		tpl.ExecuteTemplate(w, "index.html", Info.Artists)

	}

}

func bandInfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			tpl.ExecuteTemplate(w, "error.html", nil)
			fmt.Fprintln(w, " An internal server error has occurred: ", http.StatusInternalServerError)
			return

		}
	}()
	if r.URL.Path != "/bandinfo" {
		tpl.ExecuteTemplate(w, "error.html", nil)
		fmt.Fprintln(w, "\n\n\n\n Address not found: wrong address entered, status: ", http.StatusNotFound)

	} else {
		r.ParseForm()
		submit := r.FormValue("ChosenBandInfo")
		Numsubmit, _ := strconv.Atoi(submit)
		fmt.Println(Numsubmit - 1)

		fmt.Fprintln(w, "<h1>"+Info.Artists[Numsubmit-1].Name+"</h1>")

		fmt.Fprintln(w, " <h1> MEMBERS </h1> ")
		fmt.Fprintln(w, "<pre>"+strings.Join(Info.Artists[Numsubmit-1].Members, "\n"))

		fmt.Fprintln(w, "<h1> CREATION DATE </h1>")
		fmt.Fprintln(w, Info.Artists[Numsubmit-1].CreationDate)

		fmt.Fprintln(w, "<h1> FIRST ALBUM </h1>")
		fmt.Fprintln(w, Info.Artists[Numsubmit-1].FirstAlbum)

		fmt.Fprintln(w, "<h1> LOCATIONS </h1>")
		fmt.Fprintln(w, "<pre>"+strings.Join(Info.Locations[Numsubmit-1].Locations, "\n"))

		fmt.Fprintln(w, " <h1> DATES </h1> ")
		fmt.Fprintln(w, "<pre>"+strings.Join(Info.Dates[Numsubmit-1].Dates, "\n"))

		fmt.Fprintln(w, "<h1> DATES & LOCATIONS </h1>")
		for _, v := range Info.Locations[Numsubmit-1].Locations {
			fmt.Fprintln(w, "<h2>"+"<pre>"+v+"</pre>"+"</h2>")
			for _, x := range Info.Relations {

				for _, q := range x.DatesLocations[v] {
					fmt.Fprintln(w, q)
				}
			}
			fmt.Fprintln(w, "<br>")
		}

		fmt.Println(Numsubmit - 1)

		tpl.ExecuteTemplate(w, "bandinfo.html", nil)

	}
}
