package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/amosehiguese/subdscanner/scanners"
	"github.com/gorilla/mux"
)

func GetDomain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	domain := params["domain"]
	_, err := url.Parse(domain)
	if err != nil {
		log.Println("Unable to parse ->", domain)
		errResp := NewError(http.StatusBadRequest, "bad request")
		json.NewEncoder(w).Encode(errResp)
		return
	}

	subdomains, err := scanners.Scan(domain)
	if err != nil {
		log.Println("An Error occurred during scanning")
		errResp := NewError(http.StatusInternalServerError, "An error occurred. We are on it!")
		json.NewEncoder(w).Encode(errResp)
	}

	var response Subdomains
	for s := range subdomains{
		response.Subdomains = append(response.Subdomains, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
