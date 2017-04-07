package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/urfave/negroni"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	w.Write([]byte("hello！ server is running!!!"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)

	n.Run(":8080")
}
