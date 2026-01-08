package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
}

var db *sql.DB

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

func setupDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить данные из .env файла")
	}

	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		log.Fatal("Переменная окружения DB_DSN не задана")
	}

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Успешное подключение к базе данных!")
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupDatabase()
	defer db.Close()
	handleFunc()
}
