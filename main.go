package main

import (
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to fake-server"))
}

func superAdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Define endpoints and their responses"))
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/superAdmin", superAdminHandler)
	return mux
}

func main() {
	mux := createMux()
	log.Fatal(http.ListenAndServe(":8284", mux))
}
