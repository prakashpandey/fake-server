package main

import "net/http"

type endpoint struct {
	method  string
	handler endpointHandler
}

type endpointHandler struct {
	response string
	fn       http.HandlerFunc
}

var endpoints = map[string]endpoint{}

type addEndpointRequest struct {
	Path     string `json:"path"`
	Method   string `json:"method"`
	Response string `json:"response"`
}

func (e *addEndpointRequest) String() string {
	return e.Path + ", " + e.Method + ", " + e.Response
}

func (e *addEndpointRequest) addEndpoint() {
	endpoints[e.Path] = endpoint{
		method: e.Method,
		handler: endpointHandler{
			fn: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(e.Response))
			},
		},
	}
}
