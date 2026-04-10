package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Close()

	var posts = []Article{}
	for res.Next() {
		var post Article
		err := res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			log.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "create", nil)
}

func show_post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func contact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/contact.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "contact", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Пожалуйста, заполните все поля")
		return
	}

	stmt, err := db.Prepare("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES(?, ?, ?)")
	if err != nil {
		log.Println("Ошибка подготовки запроса:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, anons, full_text)
	if err != nil {
		log.Println("Ошибка выполнения запроса:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func delete_article(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 1. Если это GET запрос — показываем страницу подтверждения
	// if r.Method == "GET" {
	// 	var article Article
	// 	// Ищем статью, чтобы показать её заголовок пользователю перед удалением
	// 	err := db.QueryRow("SELECT id, title, anons, full_text FROM articles WHERE id = ?", id).Scan(&article.Id, &article.Title, &article.Anons, &article.FullText)
	// 	if err != nil {
	// 		http.NotFound(w, r)
	// 		return
	// 	}

	// 	t, err := template.ParseFiles("templates/delete.html", "templates/header.html", "templates/footer.html")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	t.ExecuteTemplate(w, "delete", article)
	// 	return
	// }

	if r.Method == "POST" {
		log.Println("Deleting article with id:", id)

		stmt, err := db.Prepare("DELETE FROM `articles` WHERE `id` = ?")
		if err != nil {
			log.Println("Ошибка подготовки запроса:", err)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(id)
		if err != nil {
			log.Println("Ошибка выполнения запроса:", err)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			return
		}

		log.Println("Article deleted successfully")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}