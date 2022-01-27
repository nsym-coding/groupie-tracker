package main

import ("fmt"
		"encoding/json"
		"html/template"
		"log"
		"os"
		"io/ioutil"
		"net/http"
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
		fmt.Println(responseObject[23])

	// 	 for i := 0; i < len(responseObject); i++ {
    //    	 fmt.Println(responseObject[i].BandName())
    //  }
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
	
		var responseObject Locations
	
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
		fmt.Println(err)
			}
		//fmt.Printf("%+v\n", responseObject)
		fmt.Println(responseObject.Index[23])

	// 	 for i := 0; i < len(responseObject.Index); i++ {
    //    	 fmt.Println(responseObject.Index[i])
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
		fmt.Println(responseObject.Index[23])

	// 	 for i := 0; i < len(responseObject.Index); i++ {
    //    	 fmt.Println(responseObject.Index[i])
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
		fmt.Println(responseObject.Index[23])

	// 	 for i := 0; i < len(responseObject.Index); i++ {
    //    	 fmt.Println(responseObject.Index[i])
    //  }
	}

//Handler function for the index
func home(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.Error(w, "404 address not found: wrong address entered!", http.StatusNotFound)
		} else {

			//a:= "https://groupietrackers.herokuapp.com/api/images/queen.jpeg"
			
			tpl.ExecuteTemplate(w, "home.html", nil)
		}
	}

func requests() {
		fs := http.FileServer(http.Dir("./templates"))
	
		http.Handle("/", fs)
		http.HandleFunc("/home.html", home)
		http.ListenAndServe(":8080", nil)
}

func main (){
	requests()
	artists()
	locations()
	dates()
	relation()
}