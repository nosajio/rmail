package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")

	// Create the main router and request handlers
	router := mux.NewRouter()
	router.HandleFunc("/message", HandlePostMessage).Methods("POST")

	// Apply middleware and listen for new requests
	fmt.Printf("Listening on :%v\n", port)
	http.ListenAndServe(":"+port, corsRouter(router))
}

func corsRouter(h http.Handler) http.Handler {
	corsMethods := handlers.AllowedMethods([]string{"HEAD", "POST", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsEnv := strings.Split(strings.Replace(os.Getenv("ALLOWED_ORIGINS"), " ", "", -1), ",") // -1 means 'replace all'
	corsOrigins := handlers.AllowedOrigins(originsEnv)
	return handlers.CORS(corsOrigins, corsHeaders, corsMethods)(h)
}
