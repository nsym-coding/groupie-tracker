package main

import ("fmt"
		"encoding/json"
		"log"
		"os"
		"io/ioutil"
		"net/http"
)

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

func (a Artists) BandName() (string, []string, int) {
	return a.Name, a.Members, a.CreationDate

}

type Locations struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
}

type IndexLocations struct {
	Index []Locations `json:"index"`
}

type Dates struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
}

type IndexDates struct {
	Index []Dates `json:"index"`
}

type Relations struct {
		ID int 			`json:"id"`
		DatesLocations map[string][]string 	`json:"DatesLocations`
}

type IndexRelations struct {
	Index []Relations `json:"index"`
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
		//fmt.Println(responseObject)

		 for i := 0; i < len(responseObject); i++ {
       	 fmt.Println(responseObject[i].BandName())
     }
	}
