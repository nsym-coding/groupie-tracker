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
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
} `json:"index"`
	}

// func (l Locations) LocationSpot() (int, []string){
// 	return l.ID, l.Locations
// }

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
} `json:"index"`
	}

type Relations struct {
	Index []struct {
		ID int 			`json:"id"`
		DatesLocations map[string][]string 	`json:"DatesLocations`
} `json:"index`
	}


// func (r Relations) RelationData() (int, map[string][]string) {
// 	return r.ID, r.DatesLocations

// }

func artists() {
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

func locations() {
		response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(string(responseData))
	
		var responseObject []Locations
	
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
		fmt.Println(err)
			}
		//fmt.Printf("%+v\n", responseObject)
		//fmt.Println(responseObject)

	// 	 for i := 0; i < len(responseObject); i++ {
    //    	 fmt.Println(responseObject[i].LocationSpot())
    //  }
	}

func dates(){
		response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(string(responseData))
	
		var responseObject Dates
	
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
		fmt.Println(err)
			}
		//fmt.Printf("%+v\n", responseObject)
		fmt.Println(responseObject)

	// 	 for i := 0; i < len(responseObject); i++ {
    //    	 fmt.Println(responseObject[i].BandName())
    //  }
	}

func relation() {


		response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(string(responseData))
	
		var responseObject Relations
	
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
		fmt.Println(err)
			}
		//fmt.Printf("%+v\n", responseObject)
		fmt.Println(responseObject)

	// 	 for i := 0; i < len(responseObject); i++ {
    //    	 fmt.Println(responseObject[i].RelationData())
    //  }
	}



func main (){
	artists()
	//locations()
	//dates()
	//relation()
}