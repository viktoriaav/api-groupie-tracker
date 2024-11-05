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

type Info struct {
	Id           int                 `json:"id"`
	Name         string              `json:"name"`
	Image        string              `json:"image"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relations    map[string][]string `json:"relations"`
}

var artistObject []Info
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	artists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	artistData, err := ioutil.ReadAll(artists.Body)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	json.Unmarshal(artistData, &artistObject)
	type ArtistMapRelation struct {
		Id           string              `json:"id"`
		LocationDate map[string][]string `json:"datesLocations"`
	}
	for x, artist := range artistObject {
		var ArtistMapRelationData ArtistMapRelation
		ArtistMap, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + fmt.Sprint(artist.Id))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		ArtistMapData, err := ioutil.ReadAll(ArtistMap.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(ArtistMapData, &ArtistMapRelationData)
		artist.Relations = ArtistMapRelationData.LocationDate
		artistObject[x] = artist
	}
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/error404", errorPage)
	fmt.Printf("Listening on port %v\n", port)
	fmt.Println("server started . . .")
	fmt.Println("ctrl(cmd) + click: http://localhost:8080/")
	http.ListenAndServe(":"+port, mux)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // error if wrong path
		http.Redirect(w, r, "error404", http.StatusSeeOther)
		return
	}
	err := tmpl.ExecuteTemplate(w, "index", artistObject)
	if err != nil { //error if template not exists
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
	}
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl, err := template.ParseFiles("templates/error404.html")
	if err != nil {
		panic(err)
	}
	r.ParseForm()
	tmpl.Execute(w, "")
}
