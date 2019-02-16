package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/mail", HandlePostMail).Methods("POST")

	fmt.Printf("Listening on :%v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
