package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
	})

	http.ListenAndServe(":9091", proxy)

}
