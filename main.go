package main

import (
	"fmt"

	"git.learn.01founders.co/nsym_coding/groupie-tracker.git/groupie"
)

/*This var is a pointer towards template.Template that is a
pointer to help process the html.*/
// var tpl *template.Template

// /*This init function, once it's initialised, makes it so that each html file
// in the templates folder is parsed i.e. they all get looked through once and
// then stored in the memory ready to go when needed*/
// func init() {
// 	tpl = template.Must(template.ParseGlob("templates/*html"))
// }


func main() {

	groupie.UnmarshalArtistData()

	//for i := 0; i < len(groupie.ArtistID); i++ {
	fmt.Println(groupie.ArtistID[0])
	fmt.Println(groupie.ArtistImage[0])
	fmt.Println(groupie.ArtistName[0])
	fmt.Println(groupie.ArtistMembers[0])
	fmt.Println(groupie.ArtistCreationDate[0])
	fmt.Println(groupie.ArtistFirstAlbum[0])
	fmt.Println(groupie.ArtistLocations[0])
	fmt.Println(groupie.ArtistConcertDates[0])
	//}
}

// func requests() {

// 	http.HandleFunc("/", index)
// 	http.HandleFunc("/info", artistInfo)
// 	http.ListenAndServe(":8080", nil)
// 	log.Println("Server started on: http://localhost:8080")
// }
// func index(w http.ResponseWriter, r *http.Request) {


// 	//-------------Create a struct to hold unmarshalled data-----------

// 	// var TotalInfo []struct {
// 	// 	responseObjectArtists   Artists
// 	// 	responseObjectLocations Locations.Index
// 	// 	responseObjectDates     Dates
// 	// }

// 	if r.URL.Path != "/" {
// 		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
// 	} else {

// 		tpl.ExecuteTemplate(w, "index.html", groupie.ArtistsDatesLocations)
// 	}
// }

// func artistInfo(w http.ResponseWriter, r *http.Request) {

// 	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
// 	if err != nil {
// 		panic("Couldn't get the relations data!")
// 	}

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		panic("Couldn't read data for the Artists")
// 	}

// 	var responseObject Relation

// 	json.Unmarshal(responseData, &responseObject)

// 	if r.URL.Path != "/info" {
// 		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
// 	} else {

// 		tpl.ExecuteTemplate(w, "info.html", groupie.ArtistsDatesLocations)
// 	}

// }
