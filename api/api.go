package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amosehiguese/subdscanner/scanners"
	"github.com/gorilla/mux"
)

func GetDomain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := params["domain"]

	subdomains, err := scanners.Scan(url)
	if err != nil {
		log.Fatalln("An Error occurred during scanning")
		errResp := NewError(http.StatusInternalServerError, "An error occurred. We are on it!")
		json.NewEncoder(w).Encode(errResp)
	}

	response := Subdomains{Subdomains: subdomains}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
