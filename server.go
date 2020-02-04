package main

import (
	"log"
	"net/http"

	"./controllers/translate"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handle)
	// De morse a humano
	router.HandleFunc("/2text", translate.HandleTranslateHuman).Methods("POST")
	// De humano a morse
	router.HandleFunc("/2morse", translate.HandleTranslateMorse).Methods("POST")
	// De bits a morse
	router.HandleFunc("/2bitsmorse", translate.HandleTranslateBitsMorse).Methods("POST")
	// De bits a humano
	router.HandleFunc("/2bitstext", translate.HandleTranslateBitsHuman).Methods("POST")
	// De morse a bits
	router.HandleFunc("/2morsebits", translate.HandleTranslateMorseBits).Methods("POST")
	// De humano a bits
	router.HandleFunc("/2humanbits", translate.HandleTranslateHumanBits).Methods("POST")

	http.Handle("/", router)

	log.Print("Listening on: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to TeleGraph\n"))
}
