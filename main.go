package main

import (
	"html/template"
	"log"
	"net/http"

	"git.learn.01founders.co/nsym_coding/groupie-tracker.git/groupie"
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

func main() {
	groupie.UnmarshalArtistData()
	groupie.UnmarshalDatesData()
	groupie.UnmarshallLocationsData()

	requests()
}

func requests() {
	http.HandleFunc("/", index)
	http.HandleFunc("/info", artistInfo)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started on: http://localhost:8080")
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {
		tpl.ExecuteTemplate(w, "index.html", groupie.Info.Artists)
	}
}

func artistInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info" {
		http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
	} else {
		tpl.ExecuteTemplate(w, "info.html", groupie.Info.Relations)
	}
}
