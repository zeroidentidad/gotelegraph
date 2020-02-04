package translate

import (
	"encoding/json"
	"log"
	"net/http"

	"../../telegraph"
)

type trasnlate struct {
	Text string `json:"text"`
}

type response struct {
	Code     int
	Response string
}

// HandleTranslateHuman maneja la traduccion de morse a humano
func HandleTranslateHuman(w http.ResponseWriter, r *http.Request) {
	var trans trasnlate
	var res response

	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		log.Print(err.Error())
	}

	res.Response = telegraph.Translate2Human(trans.Text)
	res.Code = http.StatusOK

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

// HandleTranslateMorse maneja la traduccion de texto humano a codigo morse
func HandleTranslateMorse(w http.ResponseWriter, r *http.Request) {
	var trans trasnlate
	var res response

	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		log.Print(err.Error())
	}

	res.Response = telegraph.Translate2Morse(trans.Text)
	res.Code = http.StatusOK

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

// HandleTranslateBitsMorse maneja la traduccion de bits a codigo morse
func HandleTranslateBitsMorse(w http.ResponseWriter, r *http.Request) {
	var trans trasnlate
	var res response

	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		log.Print(err.Error())
	}

	res.Response = telegraph.DecodeBits2Morse(trans.Text)
	res.Code = http.StatusOK

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

// HandleTranslateBitsHuman maneja la traduccion de bits a texto
func HandleTranslateBitsHuman(w http.ResponseWriter, r *http.Request) {
	var trans trasnlate
	var res response

	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		log.Print(err.Error())
	}

	morseMessage := telegraph.DecodeBits2Morse(trans.Text)

	res.Response = telegraph.Translate2Human(morseMessage)
	res.Code = http.StatusOK

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

// HandleTranslateMorseBits maneja la traduccion de codigo morse a bits
func HandleTranslateMorseBits(w http.ResponseWriter, r *http.Request) {
	var trans trasnlate
	var res response

	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		log.Print(err.Error())
	}

	res.Response = telegraph.DecodeMorse2Bits(trans.Text)
	res.Code = http.StatusOK

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

// HandleTranslateHumanBits maneja la traduccion de texto a bits
func HandleTranslateHumanBits(w http.ResponseWriter, r *http.Request) {
	var trans trasnlate
	var res response

	err := json.NewDecoder(r.Body).Decode(&trans)
	if err != nil {
		log.Print(err.Error())
	}

	morseMessage := telegraph.Translate2Morse(trans.Text)

	res.Response = telegraph.DecodeMorse2Bits(morseMessage)
	res.Code = http.StatusOK

	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}
