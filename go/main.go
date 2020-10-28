package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type mytype map[string]string

func main() {
	// http.HandleFunc("/", sayhello)         // Устанавливаем роутер
	http.HandleFunc("/", subcsribe)

	err := http.ListenAndServe(":80", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
	log.Println("Привет!")
}

func subcsribe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var str mytype
		err = json.Unmarshal(b, &str)
		if err != nil {
			log.Fatal("Error Decoding: ", err)
		}

	log.Println(str["url"])
	log.Println(str["email"])
	fmt.Fprintf(w, "Вы будете получать обновления")

	default:
		fmt.Println("Only POST support")
		log.Println("Only POST support")
	}

}