package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

var jsonData []Artists

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

func (a Artists) ArtistInfo() string {
	return a.Image

}

var sArtists []Artists

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

func unmarshalRelations() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := Relation{}
	// fmt.Printf("this is type %T \n", jsonData)

	err = json.Unmarshal(responseData, &jsonData)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(jsonData.Index); i++ {
		fmt.Println(jsonData.Index[i])

	}
	//fmt.Println(jsonData)
}

func unmarshalArtists() {

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := []Artists{}

	err = json.Unmarshal(responseData, &jsonData)

	if err != nil {
		panic(err)
	}

	//fmt.Println(jsonData)

	fmt.Println(jsonData[0:2])

	//a := Artists{Name: `json:"name"`}

	// for i := 0; i < len(jsonData); i++ {
	// 	fmt.Println(jsonData[i].ArtistInfo())

	// }

}

func unmarshalDates() {

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := Dates{}

	err = json.Unmarshal(responseData, &jsonData)

	if err != nil {
		panic(err)
	}

	fmt.Println(jsonData)

	//fmt.Println(jsonData.BandName)

	//a := Artists{Name: `json:"name"`}

	// for i := 0; i < len(jsonData.Index); i++ {
	// 	fmt.Println(jsonData.Index[i])

	// }

}

func unmarshalLocations() {

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := Locations{}

	err = json.Unmarshal(responseData, &jsonData)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("type of this %T:", jsonData)

	//fmt.Println(jsonData.Index[1])

	//fmt.Println(jsonData.BandName)

	//a := Artists{Name: `json:"name"`}

	for i := 0; i < len(jsonData.Index); i++ {
		fmt.Println(jsonData.Index[i])

	}

}

func requests() {
	fs := http.FileServer(http.Dir("./templates"))

	http.ListenAndServe(":8080", nil)
	http.Handle("/", fs)
	http.HandleFunc("/index.html", index)
}

func main() {
	//unmarshalRelations()
	//unmarshalLocations()
	//unmarshalDates()
	//unmarshalArtists()

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := []Artists{}

	err = json.Unmarshal(responseData, &jsonData)

	if err != nil {
		panic(err)
	}

	//fmt.Println(jsonData)

	fmt.Println(jsonData[0:2])

	fmt.Println()

	fs := http.FileServer(http.Dir("./templates"))

	http.Handle("/", fs)
	http.HandleFunc("/index.html", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(jsonData)
	if r.URL.Path != "/" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {

		tpl.ExecuteTemplate(w, "index.html", jsonData[0])
	}
}
