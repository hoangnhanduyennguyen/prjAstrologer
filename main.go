package main

import (
	"encoding/json"
	"example/prjAstrology/facts"
	"example/prjAstrology/numerology"
	"example/prjAstrology/zodiac"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func getFunFacts(w http.ResponseWriter, r *http.Request) {
	dob := r.URL.Query().Get("dob")
	result := facts.GetFunFacts(dob)

	//Marshal returns the JSON encoding
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func getZodiac(w http.ResponseWriter, r *http.Request) {
	dob := r.URL.Query().Get("dob")
	result := zodiac.GetZodiac(dob)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func getNumerology(w http.ResponseWriter, r *http.Request) {
	dob := r.URL.Query().Get("dob")
	firstName := r.URL.Query().Get("firstName")
	lastName := r.URL.Query().Get("lastName")
	result := numerology.GetNumerology(dob, firstName, lastName)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/", homePage) //HandleFunc registers the handler function for the given pattern
	http.HandleFunc("/facts", getFunFacts)
	http.HandleFunc("/zodiac", getZodiac)
	http.HandleFunc("/numerology", getNumerology)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil) //initialize http server, this function launch go default http server on our desire port.
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
