package main

import (
	"git.learn.01founders.co/nsym_coding/groupie-tracker.git/groupie"
)


func main() {

groupie.UnmarshalArtistData()

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

		tpl.ExecuteTemplate(w, "info.html", responseObject.Index)
	}

}