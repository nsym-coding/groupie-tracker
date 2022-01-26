package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	}
}

func main() {

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

	var jsonData Relation

	err = json.Unmarshal(responseData, &jsonData)

	if err != nil {
		panic(err)
	}

	fmt.Println(jsonData)
	// for i := 0; i < len(jsonData); i++ {
	// 	fmt.Println(jsonData[i])

	// }

	//fmt.Println(string(responseData))

	// var responseObject Config
	// json.Unmarshal(responseData, &responseObject)

	// fmt.Println(len(responseObject.Artists.Name))

	//fmt.Printf(responseObject.Artists.FirstAlbum)

	// for _, loc := range responseObject.Artists.Name {
	// 	fmt.Println(loc)

}
