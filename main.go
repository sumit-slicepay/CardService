package main

import (
	"log"
	"net/http"

	"controller"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	cardController := controller.NewCardController()

	router.HandleFunc("/getCVV/{uuid}", cardController.GetCVV).Methods("GET")
	router.HandleFunc("/ping", cardController.Ping).Methods("GET")

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
