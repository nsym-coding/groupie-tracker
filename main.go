package main

import ("fmt"
		"encoding/json"
		"log"
		"os"
		"io/ioutil"
		"net/http"
)

// type Groupie struct{ 

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

type Relations struct {
	Index []struct{
		ID int 			`json:"id"`
		DatesLocations map[string][]string 	`json:"DatesLocations`
	}
}

// }
func main() {
		response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(string(responseData))
	
		var responseObject []Artists
	
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
		fmt.Println(err)
			}
		//fmt.Printf("%+v\n", responseObject)
		fmt.Println(responseObject)

	// 	 for i := 0; i < len(responseObject.Artists); i++ {
    //    	 fmt.Println(responseObject.Artists[i])
    //  }
	}
