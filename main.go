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

func playRound(writer http.ResponseWriter, request *http.Request) {
	playerChoice, _ := strconv.Atoi(request.URL.Query().Get("c")) // toma el valor que clickeo el usuario
	result := rpslv.PlayRound(playerChoice)                       // con base a eso saca el resultado
	out, err := json.MarshalIndent(result, "", "       ")         // out trae la info de rpslv.go
	if err != nil {
		log.Println(err)
		return
	}
	writer.Header().Set("Content-Type", "application/json") // de esta forma se parsea el resultado a json
	writer.Write(out)                                       // lo que termina siendo el resultado
}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage) // es la request al home o la pagina raiz

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(writer http.ResponseWriter, page string) {
	temp, err := template.ParseFiles(page) // temp de template
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
