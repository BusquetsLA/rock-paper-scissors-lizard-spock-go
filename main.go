package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpslv/rpslv"
	"strconv"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "index.html")
}

// playRound is a function that handles a round of the rock-paper-scissors-lizard-spock game.
// It takes a writer http.ResponseWriter and a request *http.Request as parameters.
func playRound(writer http.ResponseWriter, request *http.Request) {
	playerChoice, _ := strconv.Atoi(request.URL.Query().Get("c")) // The player's choice is retrieved from the request URL query parameter "c".
	result := rpslv.PlayRound(playerChoice)                       // The function calls the PlayRound function from rpslv package to determine the result of the round.
	out, err := json.MarshalIndent(result, "", "       ")         // The result is then marshaled into JSON format using json.MarshalIndent.

	if err != nil {
		log.Println(err)
		return
	}
	writer.Header().Set("Content-Type", "application/json") // The result is parsed as JSON.
	writer.Write(out)                                       // The marshaled result is written to the response writer.

}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage) // The homePage function is registered to handle the root URL path.

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(writer http.ResponseWriter, page string) {
	temp, err := template.ParseFiles(page) // The template.ParseFiles function is used to parse the HTML template file.
	if err != nil {
		log.Println(err)
		return
	}
	err = temp.Execute(writer, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
