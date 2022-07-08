package main

import (
	"fmt"
	"log"
	"net/http"

	"main/api/bountie"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerClients struct {
	bountie bountie.Bountie
}

func handleRequests(clients handlerClients) {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/decimal-binary", clients.bountie.DecimalToBinary).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/longest-palindromic-string", clients.bountie.LongestPalindromicString).Methods(http.MethodPost, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	validator := validator.New()

	bountieClient := bountie.NewBountieClient(validator)

	clients := handlerClients{
		bountie: bountieClient,
	}

	fmt.Println("Listening at *:5000")
	handleRequests(clients)

}
