package main

import (
	"net/http"
) 

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(res, req, "main.wasm")
	})

	http.ListenAndServe(":8080", nil)

}
