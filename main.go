package main

import ("fmt"
		"encoding/json"
)

type Groupie struct{ 

Artist struct {
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
}

func main() {
		response, err := http.Get("https://groupietrackers.herokuapp.com/api")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
	
		var responseObject Response
		json.Unmarshal(responseData, &responseObject)
	
		// fmt.Println(responseObject.Name)
		// fmt.Println(len(responseObject.Pokemon))
	
		// // for i := 0; i < len(responseObject.Pokemon); i++ {
		// //     fmt.Println(responseObject.Pokemon[i].Species.Name)
		// // }
	
	}
