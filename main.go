package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health/", handleHealth)
	log.Println("Starting http server ...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Print("Failed to start http listener")
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"status\":\"OK\"}"))
}
