package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleFunc() {
	rt := mux.NewRouter()

	rt.HandleFunc("/", index).Methods("GET")
	rt.HandleFunc("/create", create).Methods("GET")
	rt.HandleFunc("/save_article", save_article).Methods("POST")
	rt.HandleFunc("/delete/{id:[0-9]+}", delete_article).Methods("POST")
	rt.HandleFunc("/contact", contact).Methods("GET")
	rt.HandleFunc("/articles/{id:[0-9]+}", show_post).Methods("GET")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", rt)
	http.ListenAndServe(":8080", nil)
}
