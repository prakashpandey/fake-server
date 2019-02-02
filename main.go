package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if endpoint, ok := endpoints[r.URL.Path]; !ok {
		w.Write([]byte("Endpoint is not defined"))
		return
	} else {
		if endpoint.method != r.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		} else {
			endpoint.handler.fn(w, r)
		}
	}

}

func superAdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Define endpoints and their responses"))
	} else if r.Method == http.MethodPost {
		e := &addEndpointRequest{}
		b, _ := ioutil.ReadAll(r.Body)
		log.Printf("Bytes received: %s", string(b))
		_ = json.Unmarshal(b, e)
		e.addEndpoint()
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Endpoint added"))
	}
}

func getEndpoints(w http.ResponseWriter, r *http.Request) {
	ep := make([]string, 1, 1)
	for k := range endpoints {
		ep = append(ep, k)
	}
	b, _ := json.Marshal(ep)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/superAdmin", superAdminHandler)
	mux.HandleFunc("/getEndpoints", getEndpoints)
	return mux
}

func main() {
	mux := createMux()
	log.Fatal(http.ListenAndServe(":8284", mux))
}
