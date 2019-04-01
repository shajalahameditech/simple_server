package main

import (
	"fmt"
	"net/http"
)

const DEFAULT_PORT  = 8080

func handle(w http.ResponseWriter,r *http.Request ){
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main(){
	var port int = DEFAULT_PORT


	fmt.Printf("Running on :%d", port)

	http.HandleFunc("/",handle)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
