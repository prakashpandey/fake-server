package main

import (
	"log"
	"net/http"
)

type endpoint map[string]http.HandlerFunc

var fooEndpoint = endpoint{
	"GET": func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You called Get Method of foo handler"))
	},
}

var endpoints = map[string]endpoint{
	"/foo": fooEndpoint,
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if methods, ok := endpoints[r.URL.Path]; !ok {
		w.Write([]byte("Endpoint is not defined"))
		return
	} else {
		if fn, ok := methods[r.Method]; !ok {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		} else {
			fn(w, r)
		}
	}

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
