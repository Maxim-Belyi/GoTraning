package main 

import "net/http"

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}